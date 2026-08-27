package main

import (
	"cadance/internal/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load(".env")
	if err!=nil{
		fmt.Printf("Error :%v\n",err)
		os.Exit(1)
	}

	err = server.StartBackend(os.Getenv("PORT"))
	if err!=nil{
		fmt.Printf("Error :%v\n",err)
		os.Exit(1)
	}
	
}