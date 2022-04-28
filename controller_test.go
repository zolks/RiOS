package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// Grab our router
var router *gin.Engine

//var RedisClient *redis.Client
var redisMock redismock.ClientMock

func init() {
	router = SetupRouter()
	flowMap = NewFlowMapTTL(10, 30, time.Second*10)

	//redisMock.ExpectGet("0800994455").SetVal(`{"id":1001,"service_number":"0800994455","name":"Fluxo Teste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Você ligou para a central vivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"End Error","end_cause":"504"},"103":{"id":103,"type":"End","name":"End OK","end_cause":"200"}}}`)
	//redisMock.Regexp().ExpectSet("0800994455", `[a-z]+`, 0).SetVal("OK")
	//requestData := url.Values{}
	//	requestData.Set("Ani", "21981024950")
	//	requestData.Set("Dnis", "0800994455")
}

func TestRegisterFlowValidJson(t *testing.T) {

	RedisClient, redisMock = redismock.NewClientMock()
	redisMock.Regexp().ExpectSet("0800994455", `[a-z]+`, 0).SetVal("OK")

	var fluxoJsonString string = `{"id":1001,"service_number":"0800994455","name":"Fluxo Teste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Você ligou para a central vivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"End Error","end_cause":"504"},"103":{"id":103,"type":"End","name":"End OK","end_cause":"200"}}}`

	w := performPutRequest(router, "/registerFlow", fluxoJsonString)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	message, exists := response["message"]
	assert.True(t, exists)
	assert.Equal(t, "OK", message)

}

func TestRegisterFlowInValidJsonError(t *testing.T) {

	RedisClient, redisMock = redismock.NewClientMock()
	redisMock.Regexp().ExpectSet("0800994455", `[a-z]+`, 0).SetVal("OK")

	var fluxoJsonString string = `{"id":1001,"service_number:"200"}}}`

	w := performPutRequest(router, "/registerFlow", fluxoJsonString)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	message, existsMessage := response["message"]
	assert.True(t, existsMessage)
	assert.Equal(t, "invalid character '2' after object key", message)

	status, existStatus := response["status"]
	assert.True(t, existStatus)
	assert.Equal(t, float64(-1), status)

}

func TestRegisterFlowWithoutServiceNumberError(t *testing.T) {

	RedisClient, redisMock = redismock.NewClientMock()
	redisMock.Regexp().ExpectSet("0800994455", `[a-z]+`, 0).SetVal("OK")

	var fluxoJsonString string = `{"id":1001,"service_number":"","name":"Fluxo Teste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Você ligou para a central vivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"End Error","end_cause":"504"},"103":{"id":103,"type":"End","name":"End OK","end_cause":"200"}}}`

	w := performPutRequest(router, "/registerFlow", fluxoJsonString)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	message, existsMessage := response["message"]
	assert.True(t, existsMessage)
	assert.Equal(t, "flow.service_number cannot be empty.", message)

	status, existStatus := response["status"]
	assert.True(t, existStatus)
	assert.Equal(t, float64(-1), status)

}

func TestPerformEstablishSuccess(t *testing.T) {

	RedisClient, redisMock = redismock.NewClientMock()
	redisMock.ExpectGet("0800994455").SetVal(`{"id":1001,"service_number":"0800994455","name":"Fluxo Teste","start_node":100,"nodes":{"100":{"id":100,"type":"Start","name":"Início","welcome_message":"Você ligou para a central vivo","next_node_id":101},"101":{"id":101,"type":"Callcenter","name":"CC_01","cc_number":"0800-998787","default_node_id":103,"error_node_id":102},"102":{"id":102,"type":"End","name":"End Error","end_cause":"504"},"103":{"id":103,"type":"End","name":"End OK","end_cause":"200"}}}`)

	paramData := ParamsCall{
		Ani:  "21981024950",
		Dnis: "0800994455",
	}

	jsonBytesParamData, _ := json.Marshal(paramData)
	w := performPutRequest(router, "/establish", string(jsonBytesParamData))
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	message, exists := response["message"]
	assert.True(t, exists)
	assert.Equal(t, "OK", message)

}

func performGetRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	return performRequest(r, http.MethodGet, path, nil)
}

func performGetRequestQuerystring(r http.Handler, path string, params url.Values) *httptest.ResponseRecorder {
	return performRequest(r, http.MethodGet, path, strings.NewReader(params.Encode()))
}

func performPutRequest(r http.Handler, path string, params string) *httptest.ResponseRecorder {
	return performRequest(r, http.MethodPut, path, strings.NewReader(params))
}

func performRequest(r http.Handler, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
