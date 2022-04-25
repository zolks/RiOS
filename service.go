package main

type OrchestrationService struct {
}

func (us *OrchestrationService) ExecuteEstablish(params *ParamsCall) error {

	flow := getFlowByDnis(params.Dnis)

	//TODO: validate ParamsCall fields.
	//TODO: execute all call logic
	err := flow.PerformCall(0, params);
	if err != nil {
		return err
	}

	return nil
}

func getFlowByDnis(dnis string) Flow {
	return Flow{}
}
