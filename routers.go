package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.PUT("/establish", ExecuteEstablishController)
	r.PUT("/registerFlow", RegisterFlowController)
	return r
}
