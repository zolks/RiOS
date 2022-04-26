package main

import (
	"encoding/json"
	"net/http"
)

const (
	Success                 = 0
	Error                   = -1
	InvalidEstablishRequest = -100
	InvalidFlowRequest      = -101
)

var mapStatusCode = map[int]string{
	Success:                 "OK",
	InvalidEstablishRequest: "Invalid establish request.",
	InvalidFlowRequest:      "Invalid flow register request.",
}

//Message returns map data
func Message(statusCode int) map[string]interface{} {
	return map[string]interface{}{"status": statusCode, "message": mapStatusCode[statusCode]}
}

//Message returns unrecovery error message
func ErrorMessage(err error) map[string]interface{} {
	return map[string]interface{}{"status": Error, "message": err.Error()}
}

//Respond returns basic response structure
func Respond(w http.ResponseWriter, httpStatus int, data map[string]interface{}) {
	w.WriteHeader(httpStatus)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
