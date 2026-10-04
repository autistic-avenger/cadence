package runjobs

import (
	"cadence-worker/db"
	"cadence-worker/downloader"
	"cadence-worker/picker"
	"cadence-worker/transcriber"
	"database/sql"
	"fmt"
	"os"
)

func ExecJob(dbConn *sql.DB, jobID string) {
	fmt.Println("[GETTING URL] :",jobID)
	url ,err := db.GetVideoUrl(jobID,dbConn)
	if err!=nil{
		fmt.Println("Error Getting Video URL !")
		os.Exit(1)
	}


	db.AlterJobStatus(dbConn,jobID,"downloading")
	fmt.Println("[DOWNLOADING] :",jobID)

	err = downloader.DownloadVideo(url,jobID)
	if err!=nil{
		fmt.Println("[ERROR]: ",err)
		err = db.AlterJobStatus(dbConn,jobID,"failed")
		if err!=nil{
			fmt.Println("Error Altering JOB Status")
		}
		return
	}
	fmt.Println("[DOWNLOADING DONE] :",jobID)

	db.AlterJobStatus(dbConn,jobID,"transcribing")
	fmt.Println("[TRANSCRIBING] :",jobID)

	err = transcriber.TransCribe(jobID,url)
	if err!=nil{
		fmt.Println("[ERROR TRANSCRIBING]: ",err)
		db.AlterJobStatus(dbConn,jobID,"failed")
		return
	}

	fmt.Println("[TRANSCRIBING DONE] :",jobID)

	fmt.Println("[PICKING] :",jobID)
	db.AlterJobStatus(dbConn,jobID,"picking")

	clips ,err := picker.PickClips(jobID)
	if err!=nil{
		db.AlterJobStatus(dbConn,jobID,"failed")
		fmt.Println(err)
		return 
	}
	
	fmt.Printf("[GOT %d CLIPS] :%s\n",len(clips.Clips),jobID)
	fmt.Println("[COMPLETED] :",jobID)
	db.AlterJobStatus(dbConn,jobID,"completed")

}