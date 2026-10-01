package redishelp

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis()*redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"), 
		DB:       0,  
		Protocol: 2,
	})
	return rc
}
