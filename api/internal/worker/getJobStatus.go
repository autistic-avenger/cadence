package worker

import (
	"cadance/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	jobID := c.Query("jobid")

	dbConn,err := db.ConnectDB()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"jobStaats: error connecting to db!",
		})
		return
	}
	
	jobStatus,err := db.GetStatus(dbConn,jobID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status":jobStatus,
	})
}