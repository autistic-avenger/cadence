package main

import (
	"cadence-worker/redishelp"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err!=nil{
		fmt.Println("Error loading env!")
		os.Exit(1)
	}
	redisClient := redishelp.ConnectRedis()


	
	queueName := "videoJobs"
	for {
		job, _ := redisClient.RPop(context.Background(),queueName).Result()
		if job != ""{
			fmt.Println(job)
		}
		time.Sleep(1*time.Second)
	}
}