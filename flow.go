package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Node interface {
	execute(mapNodes *map[int]Node, paramsIn *ParamsCall) error
}

type Flow struct {
	Id            int          `json:"id"`
	ServiceNumber string       `json:"service_number"`
	Name          string       `json:"name"`
	StartNode     int          `json:"start_node"`
	Nodes         map[int]Node `json:"nodes"`
}

type ParamsCall struct {
	Ani  string `json:"ani"`
	Dnis string `json:"dnis"`
	Cc   string `json:"cc"`
}

func (f Flow) PerformCall(runFromNodeId int, paramsIn *ParamsCall) error {
	if runFromNodeId != 0 {
		return f.Nodes[runFromNodeId].execute(&f.Nodes, paramsIn)
	} else {
		return f.Nodes[f.StartNode].execute(&f.Nodes, paramsIn)
	}
}

func (f Flow) Marshall() (string, error) {
	jsonBytes, err := json.Marshal(f)
	if err != nil {
		return "", err
	} else {
		return string(jsonBytes), nil
	}
}

func (f Flow) Unmarshall(jsonString string) (Flow, error) {
	byteFlow := []byte(jsonString)
	var flow Flow

	mapData := make(map[string]interface{})
	err := json.Unmarshal(byteFlow, &mapData)
	if err != nil {
		return flow, err
	}

	flow.Id = int(mapData["id"].(float64))
	flow.ServiceNumber = mapData["service_number"].(string)
	flow.Name = mapData["name"].(string)
	flow.StartNode = int(mapData["start_node"].(float64))

	value, err := UnmarshalNodes(mapData["nodes"].(map[string]interface{}), "type", map[string]reflect.Type{
		"Start":      reflect.TypeOf(Start{}),
		"End":        reflect.TypeOf(End{}),
		"Callcenter": reflect.TypeOf(Callcenter{}),
	})

	flow.Nodes = value

	return flow, nil
}

func UnmarshalNodes(data map[string]interface{}, typeJsonField string, customTypes map[string]reflect.Type) (map[int]Node, error) {
	nodes := map[int]Node{}

	for key, value := range data {

		typeName := fmt.Sprint(value.(map[string]interface{})[typeJsonField])
		var node Node
		if ty, found := customTypes[typeName]; found {
			node = reflect.New(ty).Interface().(Node)
		}

		valueBytes, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(valueBytes, &node); err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}

		nodes[id] = node
	}
	return nodes, nil
}

//start node
type Start struct {
	Id             int    `json:"id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	WelcomeMessage string `json:"welcome_message"`
	NextNodeId     int    `json:"next_node_id"`
}

func (node Start) execute(mapNodes *map[int]Node, paramsIn *ParamsCall) error {
	log.Printf("[%v] Bem-vindo: %v", node.Name, node.WelcomeMessage)
	nextNode := (*mapNodes)[node.NextNodeId]
	return nextNode.execute(mapNodes, paramsIn)
}

//end node
type End struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	EndCause string `json:"end_cause"`
}

func (node End) execute(mapNodes *map[int]Node, paramsIn *ParamsCall) error {
	log.Printf("[%v] Chamada terminada com end cause: %v", node.Name, node.EndCause)
	return nil
}

//Cc node
type Callcenter struct {
	Id               int    `json:"id"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	CallcenterNumber string `json:"cc_number"`
	DefaultNodeId    int    `json:"default_node_id"`
	ErrorNodeId      int    `json:"error_node_id"`
}

func (node Callcenter) execute(mapNodes *map[int]Node, paramsIn *ParamsCall) error {
	log.Printf("[%v] Estamos transferindo para o callcenter: %v", node.Name, node.CallcenterNumber)
	sucess := true

	var nextNode Node
	if sucess == true {
		paramsIn.Cc = node.CallcenterNumber
		nextNode = (*mapNodes)[node.DefaultNodeId]
	} else {
		nextNode = (*mapNodes)[node.ErrorNodeId]
	}

	return nextNode.execute(mapNodes, paramsIn)
}
