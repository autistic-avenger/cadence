package db

import "database/sql"

func AlterJobStatus(dbConn *sql.DB, jobID string, status string) error {
	alterQ := `UPDATE userVideos SET jobstatus = $1 WHERE jobid = $2`
	_,err := dbConn.Exec(alterQ,status,jobID)
	if err!=nil{
		return err
	}
	return nil
}