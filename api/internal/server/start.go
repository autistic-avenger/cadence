package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartBackend(port string) error {	
	r := gin.Default()	

	if port == ""{
		return fmt.Errorf("Port not Present.")
	}

	err := r.Run(":"+port)
	
	if err!= nil{
		return err
	}
	return nil
}