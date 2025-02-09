package taskengine

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/AvaProtocol/ap-avs/core/testutil"
	"github.com/AvaProtocol/ap-avs/model"
	avsproto "github.com/AvaProtocol/ap-avs/protobuf"
)

// Test make a query to a demo graphql server to ensure our node processing work
func TestGraphlQlNodeSimpleQuery(t *testing.T) {
	node := &avsproto.GraphQLQueryNode{
		Url: "https://spacex-production.up.railway.app/",
		Query: `
          query Launch {
            company {
              ceo
            }
            launches(limit: 2, sort: "launch_date_unix", order: "ASC") {
              id
              mission_name
            }
          }
		`,
	}

	nodes := []*avsproto.TaskNode{
		&avsproto.TaskNode{
			Id:   "123abc",
			Name: "graphqlQuery",
			TaskType: &avsproto.TaskNode_GraphqlQuery{
				GraphqlQuery: node,
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
			Target: "123abc",
		},
	}

	vm, err := NewVMWithData(&model.Task{
		&avsproto.Task{
			Id:      "123abc",
			Nodes:   nodes,
			Edges:   edges,
			Trigger: trigger,
		},
	}, nil, testutil.GetTestSmartWalletConfig(), nil)

	n, _ := NewGraphqlQueryProcessor(vm, node.Url)

	step, _, err := n.Execute("123abc", node)

	if err != nil {
		t.Errorf("expected rest node run succesfull but got error: %v", err)
	}

	if !step.Success {
		t.Errorf("expected rest node run succesfully but failed")
	}

	if !strings.Contains(step.Log, "Execute GraphQL gateway.thegraph.com") {
		t.Errorf("expected log contains request trace data but not found. Log data is: %s", step.Log)
	}

	if step.Error != "" {
		t.Errorf("expected log contains request trace data but found no")
	}

	var output struct {
		Company struct {
			CEO string `json:"ceo"`
		} `json:"company"`
		Launches []struct {
			ID          string `json:"id"`
			MissionName string `json:"mission_name"`
		} `json:"launches"`
	}
	err = json.Unmarshal([]byte(step.OutputData), &output)
	if err != nil {
		t.Errorf("expected the data output in json format, but failed to decode %v", err)
	}

	if len(output.Launches) != 2 {
		t.Errorf("expected 2 launches but found %d", len(output.Launches))
	}
	if output.Launches[0].ID != "5eb87cd9ffd86e000604b32a" {
		t.Errorf("id doesn't match. expected %s got %s", "5eb87cd9ffd86e000604b32a", output.Launches[0].ID)
	}

	if output.Launches[0].MissionName != "FalconSat" {
		t.Errorf("name doesn't match. expected %s got %s", "FalconSat", output.Launches[0].MissionName)
	}

	if output.Launches[0].ID != "5eb87cdaffd86e000604b32b" {
		t.Errorf("id doesn't match. expected %s got %s", "5eb87cd9ffd86e000604b32a", output.Launches[0].ID)
	}

	if output.Launches[0].MissionName != "DemoSat" {
		t.Errorf("name doesn't match. expected %s got %s", "FalconSat", output.Launches[0].MissionName)
	}
}
