package runjobs

import (
	"cadence-worker/db"
	"database/sql"
	"fmt"
	"os"
)

func ExecJob(dbConn *sql.DB, jobID string) {
	url ,err := db.GetVideoUrl(jobID,dbConn)
	if err!=nil{
		fmt.Println("Error Getting Video URL !")
		os.Exit(1)
	}

	


}