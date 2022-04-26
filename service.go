package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type OrchestrationService struct {
}

var ctx = context.TODO()

func (us *OrchestrationService) ExecuteEstablish(params *ParamsCall) (err error) {

	flow, err := getFlowByDnis(params.Dnis)
	if err != nil {
		return err
	}

	//TODO: validate ParamsCall fields.
	//TODO: execute all call logic
	if err = flow.PerformCall(0, params); err != nil {
		//TODO: set error response
		log.Panic("Error on Establish:PerformCall: ", err)
	}

	return
}

func getFlowByDnis(dnis string) (flow Flow, err error) {

	log.Printf("FlowMap: Getting Flow for dnis: %v", dnis)
	flow = flowMap.Get(dnis)
	if flow.Id == 0 {
		log.Printf("FlowMap: Not Found for dnis: %v", dnis)
		log.Printf("Redis: Getting Flow for dnis: %v", dnis)
		jsonFlow, err := RedisClient.Get(ctx, dnis).Result()
		if err == redis.Nil {
			log.Printf("Redis: Not Found for dnis: %v", dnis)
			//TODO: try from microservice
			return flow, fmt.Errorf("Flow not found for dnis: %v", dnis)
		} else if err != nil {
			panic(err)
		} else {
			log.Printf("Redis: Found Flow for dnis: %v", dnis)
			flow, err = flow.Unmarshall(jsonFlow)
			if err != nil {
				log.Panic(err)
			}
			log.Printf("FlowMap: Add Flow for dnis: %v", dnis)
			flowMap.Put(flow.ServiceNumber, flow)
		}
	}
	return
}

func (us *OrchestrationService) RegistryFlow(jsonString string) (err error) {

	flow, err := Flow{}.Unmarshall(jsonString)
	if err != nil {
		return err
	}

	if flow.ServiceNumber == "" {
		return fmt.Errorf("flow.service_number cannot be empty.")
	}

	jsonFlow, err := flow.Marshall()
	if err != nil {
		return err
	}

	err = RedisClient.Set(ctx, flow.ServiceNumber, jsonFlow, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
