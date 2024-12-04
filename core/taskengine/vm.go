package taskengine

import (
	"fmt"
	"strings"
	"sync"

	"github.com/AvaProtocol/ap-avs/core/taskengine/macros"
	avsproto "github.com/AvaProtocol/ap-avs/protobuf"
	"github.com/expr-lang/expr"
)

type VMState string

const (
	VMStateInitialize = "vm_initialize"
	VMStateCompiled   = "vm_compiled"
	VMStateReady      = "vm_ready"
	VMStateExecuting  = "vm_executing"
	VMStateCompleted  = "vm_completed"

	TriggerEdge = "__TRIGGER__"
)

type Step struct {
	NodeID string
	Next   []string
}

// The VM is the core component that load the node information and execute them, yield finaly result
type VM struct {
	// Input raw task data
	// TaskID can be used to cache compile program
	TaskID    string
	TaskNodes map[string]*avsproto.TaskNode
	TaskEdges []*avsproto.TaskEdge

	// executin logs and result per plans
	ExecutionLogs []*avsproto.Execution_Step

	Status VMState
	mu     *sync.Mutex
	// internal state that is set through out program execution
	vars map[string]any

	plans            map[string]*Step
	entrypoint       string
	instructionCount int64
}

func NewVM() (*VM, error) {
	v := &VM{
		Status:           VMStateInitialize,
		mu:               &sync.Mutex{},
		instructionCount: 0,
	}

	return v, nil
}

func (v *VM) Reset() {
	v.ExecutionLogs = []*avsproto.Execution_Step{}
	v.plans = map[string]*Step{}
	v.entrypoint = ""
	v.Status = VMStateInitialize
	v.instructionCount = 0
}

func NewVMWithData(taskID string, nodes []*avsproto.TaskNode, edges []*avsproto.TaskEdge) (*VM, error) {
	v := &VM{
		Status:           VMStateInitialize,
		TaskEdges:        edges,
		TaskNodes:        make(map[string]*avsproto.TaskNode),
		plans:            make(map[string]*Step),
		mu:               &sync.Mutex{},
		instructionCount: 0,
	}

	for _, node := range nodes {
		v.TaskNodes[node.Id] = node
	}

	// TODO: load initial data based on trigger
	v.vars = macros.GetEnvs(map[string]interface{}{
		"trigger": map[string]interface{}{
			"data": map[string]interface{}{
				// "address": strings.ToLower(event.Address.Hex()),
				// "topics": godash.Map(event.Topics, func(topic common.Hash) string {
				// 	return "0x" + strings.ToLower(strings.TrimLeft(topic.String(), "0x0"))
				// }),
				// "data":    "0x" + common.Bytes2Hex(event.Data),
				// "tx_hash": event.TxHash
			},
		},
	})

	return v, nil
}

func (v *VM) CreateSandbox() error {
	return nil
}

// Compile generates an execution plan based on edge
func (v *VM) Compile() error {
	for _, edge := range v.TaskEdges {
		if strings.Contains(edge.Source, ".") {
			v.plans[edge.Source] = &Step{
				NodeID: edge.Source,
				Next:   []string{edge.Target},
			}

			continue
		}

		if _, ok := v.plans[edge.Target]; !ok {
			v.plans[edge.Target] = &Step{
				NodeID: edge.Target,
				Next:   []string{},
			}
		}

		if edge.Source != "__TRIGGER__" {
			if _, ok := v.plans[edge.Source]; ok {
				v.plans[edge.Source].Next = append(v.plans[edge.Source].Next, edge.Target)
			} else {
				v.plans[edge.Source] = &Step{
					NodeID: edge.Source,
					Next:   []string{edge.Target},
				}
			}
		} else {
			v.entrypoint = edge.Target
		}
	}

	// TODO Setup a timeout context
	// currentStep := v.plans[v.entrypoint]
	// for currentStep != nil {
	// 	node := v.TaskNodes[currentStep.NodeID]

	// 	v.instructions = append(v.instructions, &Instruction{
	// 		Op:    "Invoke",
	// 		Value: node.Id,
	// 	})

	// 	if len(currentStep.Next) == 0 {
	// 		break
	// 	}

	// 	// TODO: Support multiple next
	// 	for _, next := range currentStep.Next {
	// 		v.instructions = append(v.instructions, &Instruction{
	// 			Op:    "Invoke",
	// 			Value: next,
	// 		})
	// 	}

	// 	currentStep = v.plans[currentStep.Next[0]]
	// }
	v.Status = VMStateReady

	return nil
}

//func (v *VM) Plan(node) {
//	//  if not if or not foo
//	//  v.instruction = append(v.instruction, {
//	//      op: "invoke", nodeid
//	//  }
//
//	//  if if block
//	//  first we put the instruction at the end
//	//  currentPosition
//
//	//  v.instruction = append(v.instruction, the block
//
//}

// Run the program. The VM will only run once.
func (v *VM) Run() error {
	if v.Status != VMStateReady {
		return fmt.Errorf("VM isn't in ready state")
	}

	v.mu.Lock()
	v.Status = VMStateExecuting
	defer func() {
		v.Status = VMStateCompleted
		v.mu.Unlock()
	}()
	if len(v.plans) == 0 {
		return fmt.Errorf("internal error: not compiled")
	}

	// TODO Setup a timeout context
	currentStep := v.plans[v.entrypoint]
	for currentStep != nil {
		node := v.TaskNodes[currentStep.NodeID]

		_, err := v.executeNode(node)
		if err != nil {
			// abort execution as soon as a node raise error
			return err
		}

		if len(currentStep.Next) == 0 {
			break
		}

		// TODO: Support multiple next
		currentStep = v.plans[currentStep.Next[0]]
	}

	return nil
}

func (v *VM) executeNode(node *avsproto.TaskNode) (*avsproto.Execution_Step, error) {
	v.instructionCount += 1

	var err error
	executionLog := &avsproto.Execution_Step{
		NodeId: node.Id,
	}

	if nodeValue := node.GetRestApi(); nodeValue != nil {
		p := NewRestProrcessor()
		executionLog, err = p.Execute(node.Id, nodeValue)
		v.ExecutionLogs = append(v.ExecutionLogs, executionLog)
	} else if nodeValue := node.GetBranch(); nodeValue != nil {
		outcome := ""
		executionLog, outcome, err = v.runBranch(node.Id, nodeValue)
		v.ExecutionLogs = append(v.ExecutionLogs, executionLog)
		if err == nil && outcome != "" {
			if outcomeNodes := v.plans[outcome].Next; len(outcomeNodes) >= 0 {
				for _, nodeID := range outcomeNodes {
					// TODO: track stack too deepth and abort
					node := v.TaskNodes[nodeID]
					executionLog, err = v.executeNode(node)
					//v.ExecutionLogs = append(v.ExecutionLogs, executionLog)

				}
			}
		}
	}

	return executionLog, err
}

func (v *VM) runBranch(stepID string, node *avsproto.BranchNode) (*avsproto.Execution_Step, string, error) {
	s := &avsproto.Execution_Step{
		NodeId:  stepID,
		Result:  "",
		Log:     "",
		Error:   "",
		Success: true,
	}

	var sb strings.Builder

	sb.WriteString("Execute Branch: ")
	sb.WriteString(stepID)
	outcome := ""
	for _, statement := range node.Conditions {
		if strings.EqualFold(statement.Type, "else") {
			// Execute this directly
			outcome = fmt.Sprintf("%s.%s", stepID, statement.Id)
			sb.WriteString("\nuse else condition, follow path ")
			sb.WriteString(outcome)
			s.Log = sb.String()
			return s, outcome, nil
		}
		sb.WriteString(fmt.Sprintf("\nevaluate condition: %s expression: `%s`", statement.Id, statement.Expression))

		// now we need to valuate condition
		program, err := expr.Compile(statement.Expression, expr.Env(v.vars), expr.AsBool())
		if err != nil {
			s.Success = false
			s.Error = fmt.Errorf("error compile the statement: %w", err).Error()
			sb.WriteString("error compile expression")
			s.Log = sb.String()
			return s, outcome, fmt.Errorf("error compile the statement: %w", err)
		}
		result, err := expr.Run(program, v.vars)
		if err != nil {
			s.Success = false
			s.Error = fmt.Errorf("error run statement: %w", err).Error()
			sb.WriteString("error run expression")
			s.Log = sb.String()
			return s, outcome, fmt.Errorf("error evaluate the statement: %w", err)
		}

		if result.(bool) == true {
			outcome = fmt.Sprintf("%s.%s", stepID, statement.Id)
			sb.WriteString("\nexpression result to true. follow path ")
			sb.WriteString(outcome)
			// run the node
			s.Log = sb.String()
			return s, outcome, nil
		}
	}

	sb.WriteString("\nno condition matched. halt execution")
	s.Log = sb.String()
	return s, "", nil
}
