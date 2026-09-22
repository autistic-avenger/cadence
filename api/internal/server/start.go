package server

import (
	"cadance/internal/auth"
	"cadance/internal/helpers"
	"cadance/internal/worker"
	youtubeapi "cadance/internal/youtubeAPI"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func initRedisClient() *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"), 
		DB:       0,  
		Protocol: 2,
	})
	return rc
}

func StartBackend(port string) error {	
	RedisClient := initRedisClient()
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

	r.GET("/api/video",youtubeapi.AddVideoInfo)

	r.GET("/api/getVideos",helpers.GetVideos)

	r.POST("/api/create",worker.QueueJob(RedisClient))

	err := r.Run(":"+port)
	
	if err!= nil{
		return err
	}
	return nil
}