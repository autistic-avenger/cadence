package auth

import (
	"net/http"
	"net/url"
	"os"
	"github.com/gin-gonic/gin"
)

type UserInfo struct{
	Name string `json:"name"`
	Email string `json:"email"`
}

type AcessTokenResponse struct{
	AccessToken string  `json:"access_token"` 
	ExpiresIn 	int     `json:"expires_in"` 
	RefreshToken string `json:"refresh_token"`
	IDToken		string  `json:"id_token"`
}

var GOOGLE_AUTH_SERVER string = "https://accounts.google.com/o/oauth2/v2/auth"
var GOOGLE_TOKEN string = "https://oauth2.googleapis.com/token"
var GOOGLE_OPENID string = "https://openidconnect.googleapis.com/v1/userinfo"


func HandleAuth(c *gin.Context){
	p := url.Values{}
	p.Set("client_id",os.Getenv("GOOGLE_CLIENT_ID"))
	p.Set("response_type","code")
	p.Set("redirect_uri",os.Getenv("BACKEND_URL")+"/auth/google/callback")
	p.Set("scope","openid email profile")
	p.Set("state","2de75bd30")

	redirectURI := GOOGLE_AUTH_SERVER+"?"+p.Encode()
	c.Redirect(http.StatusTemporaryRedirect,redirectURI)
}


