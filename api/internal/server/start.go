package server

import (
	"cadance/internal/auth"
	"fmt"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartBackend(port string) error {	
	r := gin.Default()	

	if port == ""{
		return fmt.Errorf("Port not Present.")
	}
	
	r.Use(cors.New(cors.Config{
		AllowOrigins: 	  []string{os.Getenv("FRONTEND_URL")},
	    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/auth/google",auth.HandleAuth)

	r.GET("/auth/google/callback",auth.HandleCallback)

	r.GET("/auth/verify",auth.VerifyJWT)


	err := r.Run(":"+port)
	
	if err!= nil{
		return err
	}
	return nil
}