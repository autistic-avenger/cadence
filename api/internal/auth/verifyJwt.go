package auth

import (
	jwthelp "cadance/internal/jwtHelp"
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
		return
	}

	_ ,err = jwt.Parse(tokenStr,jwthelp.GetSecret)

	if err!=nil{
		c.SetCookie("token","",-1,"/",os.Getenv("DOMAIN"),false,true)
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"error parsing jwt",
		})
		return
	}

	
	c.String(http.StatusOK,"Verified")
}