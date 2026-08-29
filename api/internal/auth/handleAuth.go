package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)



func HandleAuth(c *gin.Context){
	c.Redirect(http.StatusTemporaryRedirect,"https://pornhub.com")
}