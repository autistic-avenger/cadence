package worker

import (
	"cadance/internal/db"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type postBody struct {
	Url string	`json:"url"`
	JobID string `json:"jobID"`
}

func QueueJob(RedisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context){
		var body postBody
		
		err := c.ShouldBindJSON(&body)
		if err!=nil{
			c.JSON(http.StatusBadRequest,"Fuck YOU!")
			return
		}
		
		
		if body.JobID == ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"jobID Missing!"})
			return
		}

		ctx := context.Background()

		createdUnique,err := RedisClient.SetNX(ctx,body.JobID,"aura",-1).Result()
		if !createdUnique {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"Already Queued!",	
			})
			return
		}

		queueName := "videoJobs"
		err = RedisClient.LPush(ctx,queueName,body.JobID).Err()
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"LPUSH FAILED!",	
			})
			return
		}


		dbConn, err := db.ConnectDB()
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"queueing error ,!",	
			})
			return
		}

		err = db.UpdateJobStatus(dbConn,body.JobID,"queued") 
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"Error Updating Status in DB!",	
			})
			return
		}


		c.JSON(http.StatusOK,gin.H{
			"jobID":body.JobID,
		})
	}
}
