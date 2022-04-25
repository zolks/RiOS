package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	}

	response := Message(Success)
	response["data"] = params
	//return response using api helper
	Respond(c.Writer, http.StatusOK, response)

}
