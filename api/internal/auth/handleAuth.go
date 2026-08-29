package auth

import (
	"net/http"
	"net/url"
	"os"
	"github.com/gin-gonic/gin"
)



func HandleAuth(c *gin.Context){
	p := url.Values{}
	p.Set("client_id",os.Getenv("GOOGLE_CLIENT_ID"))
	p.Set("response_type","code")
	p.Set("redirect_uri",os.Getenv("BACKEND_URL")+"/auth/google/callback")
	p.Set("scope","openid email profile")
	p.Set("state","2de75bd30")

	redirectURI := "https://accounts.google.com/o/oauth2/v2/auth"+"?"+p.Encode()
	c.Redirect(http.StatusTemporaryRedirect,redirectURI)
}


func HandleCallback(c *gin.Context){
	c.Redirect(http.StatusTemporaryRedirect,"https://pornhub.com")
}