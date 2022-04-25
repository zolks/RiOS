package main

import (
	"testing"
)

var fluxoJsonString string = `{"id":1001,"name":"Fluxo Teste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Você ligou para a central vivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"End Error","end_cause":"504"},"103":{"id":103,"type":"End","name":"End OK","end_cause":"200"}}}`

var fluxoObj = Flow{
	Id:        1001,
	Name:      "Fluxo Teste",
	StartNode: 100,
	Nodes: map[int]Node{
		100: Start{
			Id:             100,
			Type:           "Start",
			Name:           "Início",
			WelcomeMessage: "Você ligou para a central vivo",
			NextNodeId:     101,
		},
		101: Callcenter{
			Id:               101,
			Type:             "Callcenter",
			Name:             "CC_01",
			CallcenterNumber: "0800-998787",
			DefaultNodeId:    103,
			ErrorNodeId:      102,
		},
		102: End{
			Id:       102,
			Type:     "End",
			Name:     "End Error",
			EndCause: "504",
		},
		103: End{
			Id:       103,
			Type:     "End",
			Name:     "End OK",
			EndCause: "200",
		},
	},
}

func TestFlow_Marshall_Success(t *testing.T) {

	jsonString, err := fluxoObj.Marshall()
	if err != nil {
		t.Fatal(err)
	}
	if fluxoJsonString != jsonString {
		t.Error("Expected:", fluxoJsonString, "Got:", jsonString)
	}

}

func TestFlow_Unmarshall_Succes(t *testing.T) {
	jsonFlow, err := Flow{}.Unmarshall(fluxoJsonString)
	if err != nil {
		t.Fatal(err)
	}
	if len(jsonFlow.Nodes) != len(fluxoObj.Nodes) {
		t.Error("Expected:", len(fluxoObj.Nodes), "Got:", len(jsonFlow.Nodes))
	}
}

func TestFlow_PerformCall(t *testing.T) {
	paramsIn := ParamsCall{}
	fluxoObj.PerformCall(0, &paramsIn)
}
