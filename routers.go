package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.PUT("/rios/establish", ExecuteEstablishController)
	r.PUT("/rios/registerFlow", RegisterFlowController)
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return r
}
