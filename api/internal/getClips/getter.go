package getclips

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetClips(c *gin.Context) {
	jobID := c.Query("id")
	jobPath := GetJobPath()
	if jobPath==""{
		c.JSON(http.StatusBadRequest,"")
		return
	}
	clipsJSON := filepath.Join(jobPath,jobID,"clips.json")

	jsonData,err := os.ReadFile(clipsJSON)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,"")
		return 
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", jsonData)
}


func GetJobPath() string {
	path ,err := filepath.Abs("../worker/jobDataStore")
	if err!=nil{
		return ""
	}
	return path
}