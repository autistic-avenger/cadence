package worker

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type postBody struct {
	Url string	`json:"url"`
	JobID string `json:"jobID"`
}

func StartClipProcess(c *gin.Context) {
	// Add Redis ,Queues and, YT_DLP  Download mp3 ->  whisper Tiny -> Open Router
	var body postBody
	err := c.ShouldBindJSON(&body)
	if err!=nil{
		c.JSON(http.StatusBadRequest,"Fuck YOU!")
	}

	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDIS_PASS"), 
		DB:       0,  
		Protocol: 2,
	})

	if body.JobID == ""{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"jobID Missing!"})
		return
	}
	ctx := context.Background()
	err = rc.Get(ctx,body.JobID).Err()
	if err==nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"It's Already Added"})
		return
	}
	err = rc.Set(ctx,body.JobID,"Started",0).Err()
	

	c.JSON(200,body)
}