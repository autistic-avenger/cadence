package db

import "database/sql"

func UpdateJobStatus(dbConn *sql.DB,jobID string, jobStatus string) error {
	
	UpdateQ  := `UPDATE userVideos SET jobstatus = $1 WHERE jobid = $2`
	_,err := dbConn.Exec(UpdateQ,jobStatus,jobID)
	if err!=nil{
		return err
	}

	return nil
}