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
	ASCIIART := `     __        __         _                 
     \ \      / /__  _ __| | _____ _ __     
      \ \ /\ / / _ \| '__| |/ / _ \ '__|    
  _    \ V  V / (_) | |  |   <  __/ |     _ 
 (_)    \_/\_/ \___/|_|  |_|\_\___|_|    (_)
                                            `

	fmt.Println(ASCIIART)
	runjobs.FetchJob(redisClient)
}