package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ExecuteEstablishController(c *gin.Context) {
	var orchService OrchestrationService
	var params ParamsCall

	if err := json.NewDecoder(c.Request.Body).Decode(&params); err != nil {
		Respond(c.Writer, http.StatusBadRequest, Message(InvalidEstablishRequest))
		return
	}

	//call service
	if err := orchService.ExecuteEstablish(&params); err != nil {
		Respond(c.Writer, http.StatusInternalServerError, ErrorMessage(err))
		return
	}

	response := Message(Success)
	response["data"] = params
	//return response using api helper
	Respond(c.Writer, http.StatusOK, response)

}

func RegisterFlowController(c *gin.Context) {
	var orchService OrchestrationService
	var flowJsonString []byte

	flowJsonString, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Respond(c.Writer, http.StatusBadRequest, Message(InvalidEstablishRequest))
		return
	}

	if err := orchService.RegistryFlow(string(flowJsonString)); err != nil {
		Respond(c.Writer, http.StatusInternalServerError, ErrorMessage(err))
		return
	}

	Respond(c.Writer, http.StatusOK, Message(Success))
}
