package youtubeapi

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type YoutubeResponse struct{
	Title string `json:"title"`
	Author string `json:"author_name"`
	ThumbnailURI string `json:"thumbnail_url"`
}

func GetInfo(c *gin.Context) {
	vidID := c.Query("uid")
	if vidID == ""{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"No Query Attached.",
		})
		return
	}

	YOUTUBE_API_URL := "https://www.youtube.com/oembed?url=https://www.youtube.com/watch?v="+vidID

	parsedUrl,err := url.Parse(YOUTUBE_API_URL)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"No Query Attached.",
		})
		return
	}
	videoURL := parsedUrl.Query().Get("url")

	client := &http.Client{}

	request,err := http.NewRequest("GET",YOUTUBE_API_URL,nil)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Error From Youtube.",
		})
		return
	}
	res ,err := client.Do(request)
	
	var INFO YoutubeResponse

	err = json.NewDecoder(res.Body).Decode(&INFO)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Error from YT.",
		})
		return
	}
	//TODO: Add to DB

	c.JSON(http.StatusOK,gin.H{
		"title":INFO.Title,
		"thumbnail":INFO.ThumbnailURI,
		"author":INFO.Author,
		"url":videoURL,
		"creator":INFO.Author,
	})

}