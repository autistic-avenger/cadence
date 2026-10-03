package runjobs

import (
	"cadence-worker/db"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var queueName string = "videoJobs"

func FetchJob(rc *redis.Client) {
	dbConn,err := db.ConnectDB()
	if err!=nil{
		fmt.Println("Error Connecting to Db!",err)
		os.Exit(1);
	}

	for {
		job, _ := rc.RPop(context.Background(),queueName).Result()
		if job != ""{
			go ExecJob(dbConn,job)
		}
		time.Sleep(3*time.Second)
	}
}
