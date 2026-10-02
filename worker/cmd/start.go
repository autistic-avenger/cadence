package main

import (
	"cadence-worker/redishelp"
	"cadence-worker/runjobs"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err!=nil{
		fmt.Println("Error loading env!")
		os.Exit(1)
	}
	redisClient := redishelp.ConnectRedis()
	runjobs.FetchJob(redisClient)
}