package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Not LoggedIN",
		})
	}

	_ ,err = jwt.Parse(tokenStr,func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")),nil
	})

	if err!=nil{
		c.SetCookie("token","",-1,"/",os.Getenv("DOMAIN"),false,true)
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"error parsing jwt",
		})
	}

	
	c.String(http.StatusOK,"Verified")
}