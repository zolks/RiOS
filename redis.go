package main

import (
	"github.com/go-redis/redis/v8"
)

//RedisClient ...
var RedisClient *redis.Client

//InitRedis ...
func InitRedis(redisHost string, selectDB ...int) {

	//var redisHost = os.Getenv("REDIS_HOST")
	//var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisHost,
		//Password: redisPassword,
		DB: selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

}

//GetRedis ...
func GetRedis() *redis.Client {
	return RedisClient
}
