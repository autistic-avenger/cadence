package runjobs

import (
	"cadence-worker/db"
	"cadence-worker/downloader"
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


	db.AlterJobStatus(dbConn,jobID,"downloading")
	
	err = downloader.DownloadVideo(url,jobID)
	if err!=nil{
		fmt.Println("DOWNLOADING [ERROR]: ",err)
		err = db.AlterJobStatus(dbConn,jobID,"failed")
		if err!=nil{
			fmt.Println("Error Altering JOB Status")
		}
	}


}