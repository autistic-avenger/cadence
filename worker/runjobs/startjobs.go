package runjobs

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var queueName string = "videoJobs"

func FetchJob(rc *redis.Client) {
	for {
		job, _ := rc.RPop(context.Background(),queueName).Result()
		if job != ""{
			fmt.Println(job)
		}
		time.Sleep(3*time.Second)
	}

}