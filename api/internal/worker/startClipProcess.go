package worker

import (
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
		err = RedisClient.Get(ctx,body.JobID).Err()
		if err==nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"It's Already Added"})
			return
		}
		err = RedisClient.Set(ctx,body.JobID,"Started",0).Err()
		
		
		c.JSON(200,body)
	}
}
