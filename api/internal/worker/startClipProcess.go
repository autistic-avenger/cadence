package worker

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(200,body)
}