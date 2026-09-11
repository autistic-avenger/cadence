package helpers

import (
	"cadance/internal/db"
	jwthelp "cadance/internal/jwtHelp"
	youtubeapi "cadance/internal/youtubeAPI"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetVideos(c *gin.Context) {
	jwtToken, err := c.Cookie("token")
	if err!=nil{
		c.JSON(http.StatusBadRequest,nil)
	}

	token ,err := jwt.ParseWithClaims(jwtToken,&youtubeapi.EmailNameClaim{},jwthelp.GetSecret)
	if err!=nil{
		c.String(http.StatusBadRequest,"Parsing jwt error!")
	}
	
	claims := token.Claims.(*youtubeapi.EmailNameClaim)
	
	dbCursor,err := db.ConnectDB()
	if err!=nil{
		c.String(http.StatusBadRequest,"DB Connecting error!")
		return 
	}

	err = db.CreateTable(dbCursor)
	if err!=nil{
		c.String(http.StatusInternalServerError,"Error creating table!")	
		return 
	}

	vids , err := db.DbGetVideos(dbCursor,claims.Email)
	if err!=nil{
		c.String(http.StatusBadRequest,"DB Fetching videos error!")
		return
	}
	vids = Reverse(vids)

	c.JSON(http.StatusOK,vids)
}



func Reverse[T any](arr []T)[]T {
	for i:= 0 ;i<len(arr)/2; i++{
		arr[i] ,arr[len(arr)-1-i] = arr[len(arr)-1-i] ,arr[i] 
	}

	return arr 
}