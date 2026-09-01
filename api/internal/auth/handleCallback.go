package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HandleCallback(c *gin.Context) {
	code := c.Query("code")
	errors := c.Query("error")
	if errors != "" {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}
	p := url.Values{}
	p.Add("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	p.Add("client_secret", os.Getenv("GOOGLE_CLIENT_SECRET"))
	p.Add("code", code)
	p.Add("grant_type", "authorization_code")
	p.Add("redirect_uri", os.Getenv("BACKEND_URL")+"/auth/google/callback")

	getTokenURI := GOOGLE_TOKEN + "?" + p.Encode()

	fmt.Println(getTokenURI)

	client := http.Client{}
	req, err := http.NewRequest("POST", getTokenURI, nil)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}

	res, err := client.Do(req)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}
	jsonString, err := io.ReadAll(res.Body)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}

	var RESPONSE AcessTokenResponse
	err = json.NewDecoder(bytes.NewBuffer(jsonString)).Decode(&RESPONSE)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}

	token, _, err := jwt.NewParser().ParseUnverified(
		RESPONSE.IDToken,
		jwt.MapClaims{},
	)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}

	var userData UserInfo
	claimes := token.Claims.(jwt.MapClaims)
	userData.Email = claimes["email"].(string)
	userData.Name = claimes["name"].(string)

	signedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":userData.Email,
		"name":userData.Name,
	})

	jwtToken ,err := signedToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err!=nil{
		c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL")+"?error=loginFail")
	}

	c.SetCookie("token",jwtToken,34560000,"/",os.Getenv("DOMAIN"),false,true)
	c.Redirect(http.StatusTemporaryRedirect, os.Getenv("FRONTEND_URL"))
}