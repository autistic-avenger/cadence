package youtubeapi

import (
	jwthelp "cadance/internal/jwtHelp"
	"database/sql"
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
	
	_ = token.Claims.(*EmailNameClaim)

	dbCursor, err := sql.Open("postgres",os.Getenv("POSTGRES_URI"))
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error connecting to db!",
		})
		return
	}
	defer dbCursor.Close()
	jobID := uuid.New().String()

	createTable := `
	CREATE TABLE IF NOT EXISTS userVideos (
		email VARCHAR(50),
		jobid VARCHAR(36),
		title VARCHAR(100),
		thumbnail VARCHAR(70),
		url VARCHAR(70),
		creator VARCHAR(60)
	)`
	
	dbRes , err := dbCursor.Exec(createTable)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Error Making Request to db!",
		})
		fmt.Println(os.Getenv("POSTGRES_URI"))
		fmt.Println(err)
		return
	}
	fmt.Println(dbRes.RowsAffected())


	c.JSON(http.StatusOK,gin.H{
		"title":INFO.Title,
		"thumbnail":INFO.ThumbnailURI,
		"author":INFO.Author,
		"url":videoURL,
		"creator":INFO.Author,
		"jobID":jobID,
	})

}