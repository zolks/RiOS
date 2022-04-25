package main

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

var flowMap *FlowMapTTL

func main() {

	router := SetupRouter()

	//init redis connection
	InitRedis(RedisAddr, 1)

	//init flowMap
	flowMap = NewFlowMapTTL(10, 20)

	router.Run(":9080")
	router.Run(ListenAddr)
}
