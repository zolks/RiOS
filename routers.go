package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.PUT("/rios/establish", ExecuteEstablishController)
	//r.PUT("/rios/establish", func(c *gin.Context) {
	//time.Sleep(2 * time.Second)
	//c.String(http.StatusInternalServerError, "ERROR")
	//})
	r.PUT("/rios/registerFlow", RegisterFlowController)
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return r
}
