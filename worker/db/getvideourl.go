package db

import "database/sql"

func GetVideoUrl(jobID string, dbConn *sql.DB) (string, error) {
	var url string 
	getUrlQ := `SELECT url from userVideos WHERE jobid = $1`

	urlRow ,err := dbConn.Query(getUrlQ,jobID)
	if err!=nil{
		return "",err
	}

	for urlRow.Next(){
		err := urlRow.Scan(&url)
		if err!=nil{
			return "", err
		}
	}
	return url, nil
}