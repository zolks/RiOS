package main

import "time"

var flowMap *FlowMapTTL

func main() {

	router := SetupRouter()

	//init redis connection
	InitRedis(getEnv("REDIS_HOST", "localhost")+":"+getEnv("REDIS_PORT", "6379"), 1)

	//init flowMap
	flowMap = NewFlowMapTTL(10, 30, time.Second*10)

	router.Run(":" + getEnv("HTTP_PORT", "9080"))
}
