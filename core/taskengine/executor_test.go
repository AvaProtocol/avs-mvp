package taskengine

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AvaProtocol/ap-avs/core/testutil"
	"github.com/AvaProtocol/ap-avs/model"
	avsproto "github.com/AvaProtocol/ap-avs/protobuf"
	"github.com/AvaProtocol/ap-avs/storage"
)

func TestExecutorAppendLog(t *testing.T) {
	fmt.Println("Pending impl")
}

func TestExecutorRunTaskStopAndReturnErrorWhenANodeFailed(t *testing.T) {
	SetRpc(testutil.GetTestRPCURL())
	SetCache(testutil.GetDefaultCache())
	db := testutil.TestMustDB()
	defer storage.Destroy(db.(*storage.BadgerStorage))

	nodes := []*avsproto.TaskNode{
		&avsproto.TaskNode{
			Id:   "branch1",
			Name: "branch",
			TaskType: &avsproto.TaskNode_Branch{
				Branch: &avsproto.BranchNode{
					Conditions: []*avsproto.Condition{
						&avsproto.Condition{
							Id:         "a1",
							Type:       "if",
							Expression: "a >= 5",
						},
					},
				},
			},
		},
		&avsproto.TaskNode{
			Id:   "notification1",
			Name: "httpnode",
			TaskType: &avsproto.TaskNode_RestApi{
				RestApi: &avsproto.RestAPINode{
					Url:    "https://httpbin.org/post",
					Method: "POST",
					Body:   "hit=notification1",
				},
			},
		},
	}

	trigger := &avsproto.TaskTrigger{
		Id:   "triggertest",
		Name: "triggertest",
	}
	edges := []*avsproto.TaskEdge{
		&avsproto.TaskEdge{
			Id:     "e1",
			Source: trigger.Id,
			Target: "branch1",
		},
		&avsproto.TaskEdge{
			Id:     "e1",
			Source: "branch1.a1",
			Target: "notification1",
		},
	}

	task := &model.Task{
		&avsproto.Task{
			Id:      "TaskID123",
			Nodes:   nodes,
			Edges:   edges,
			Trigger: trigger,
		},
	}

	executor := NewExecutor(testutil.GetTestSmartWalletConfig(), db, testutil.GetLogger())
	execution, err := executor.RunTask(task, &QueueExecutionData{
		TriggerMetadata: testutil.GetTestEventTriggerMetadata(),
		ExecutionID:     "exec123",
	})

	if err == nil || err.Error() != "Error executing task TaskID123 error evaluating the statement: ReferenceError: a is not defined at <eval>:1:8(1)" {
		t.Errorf("expect error but got %v", err)
	}

	if execution.Success {
		t.Errorf("Expect failure status but got success")
	}

	if execution.Steps[0].OutputData != "" {
		t.Errorf("output data isn't empty")
	}

	if len(execution.Steps) != 1 {
		t.Errorf("Expect evaluate one step only but got: %d", len(execution.Steps))
	}

	if execution.Steps[0].NodeId != "branch1" {
		t.Errorf("step id doesn't match, expect branch1 but got: %s", execution.Steps[0].NodeId)
	}
}

func TestExecutorRunTaskComputeSuccessFalseWhenANodeFailed(t *testing.T) {
	// Set up a test HTTP server that returns a 503 status code
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	SetRpc(testutil.GetTestRPCURL())
	SetCache(testutil.GetDefaultCache())
	db := testutil.TestMustDB()
	defer storage.Destroy(db.(*storage.BadgerStorage))

	nodes := []*avsproto.TaskNode{
		&avsproto.TaskNode{
			Id:   "branch1",
			Name: "branch",
			TaskType: &avsproto.TaskNode_Branch{
				Branch: &avsproto.BranchNode{
					Conditions: []*avsproto.Condition{
						&avsproto.Condition{
							Id:         "condition1",
							Type:       "if",
							Expression: "true",
						},
					},
				},
			},
		},
		&avsproto.TaskNode{
			Id:   "rest1",
			Name: "httpnode",
			TaskType: &avsproto.TaskNode_RestApi{
				RestApi: &avsproto.RestAPINode{
					Url:    server.URL, // Use the test server URL
					Method: "POST",
					Body:   "hit=notification1",
				},
			},
		},
	}

	trigger := &avsproto.TaskTrigger{
		Id:   "triggertest",
		Name: "triggertest",
	}
	edges := []*avsproto.TaskEdge{
		&avsproto.TaskEdge{
			Id:     "e1",
			Source: trigger.Id,
			Target: "branch1",
		},
		&avsproto.TaskEdge{
			Id:     "e1",
			Source: "branch1.condition1",
			Target: "rest1",
		},
	}

	task := &model.Task{
		&avsproto.Task{
			Id:      "TaskID123",
			Nodes:   nodes,
			Edges:   edges,
			Trigger: trigger,
		},
	}

	executor := NewExecutor(testutil.GetTestSmartWalletConfig(), db, testutil.GetLogger())
	execution, err := executor.RunTask(task, &QueueExecutionData{
		TriggerMetadata: testutil.GetTestEventTriggerMetadata(),
		ExecutionID:     "exec123",
	})

	if err == nil {
		t.Errorf("expected error due to 503 response but got nil")
	}

	if execution.Success {
		t.Errorf("expected failure status but got success")
	}

	if len(execution.Steps) != 2 {
		t.Errorf("Expect evaluate 2 steps but got: %d", len(execution.Steps))
	}

	if execution.Steps[0].NodeId != "branch1" {
		t.Errorf("step id doesn't match, expect branch1 but got: %s", execution.Steps[0].NodeId)
	}
	if execution.Steps[1].NodeId != "rest1" {
		t.Errorf("step id doesn't match, expect branch1 but got: %s", execution.Steps[0].NodeId)
	}


	// Additional assertions can be added here based on the expected behavior
}

