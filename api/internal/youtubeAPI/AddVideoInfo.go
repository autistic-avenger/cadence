package youtubeapi

import (
	"cadance/internal/db"
	jwthelp "cadance/internal/jwtHelp"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type YoutubeResponse struct{
	Title string `json:"title"`
	Author string `json:"author_name"`
	ThumbnailURI string `json:"thumbnail_url"`
}

type EmailNameClaim struct{
	Email string `json:"email"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func AddVideoInfo(c *gin.Context) {
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
	unverifiedJWT ,err := c.Cookie("token")
	if err!=nil{
		fmt.Println("No Token Found")
		c.JSON(http.StatusBadRequest,nil)
		return 
	}

	token,err := jwt.ParseWithClaims(unverifiedJWT,&EmailNameClaim{},jwthelp.GetSecret)
	if err!=nil{
		fmt.Println("Bad Token")
		c.SetCookie("token","",-1,"/",os.Getenv("DOMAIN"),false,true)
		c.Redirect(http.StatusTemporaryRedirect,os.Getenv("FRONTEND_URL"))
		return 
	}
	
	claims := token.Claims.(*EmailNameClaim)

	
	jobID := uuid.New().String()

	dbCon ,err := db.ConnectDB()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error Connecting to db.",
		})
		return 
	}

	err = db.CreateTable(dbCon)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error Creating table",
		})
		return 
	}


	data := db.YoutubeDbInfo{
		Email: claims.Email,
		Title: INFO.Title,
		Thumbnail: INFO.ThumbnailURI,
		Url: videoURL,
		Creator:INFO.Author ,
		JobID: jobID,
	}
	err = db.AddVideoInfo(dbCon,data)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error Adding to Table.",
		})
		return 
	}



	c.JSON(http.StatusOK,gin.H{
		"email":claims.Email,
		"title":INFO.Title,
		"thumbnail":INFO.ThumbnailURI,
		"url":videoURL,
		"creator":INFO.Author,
		"jobID":jobID,
	})

}