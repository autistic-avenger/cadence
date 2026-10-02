package db

import "database/sql"

func GetStatus(dbCon *sql.DB, jobID string)(string,error){
	defer dbCon.Close()
	getStatusQ := `SELECT jobstatus from userVideos WHERE jobid = $1`

	res ,err := dbCon.Query(getStatusQ,jobID)
	if err!=nil{
		return  "",err
	}
	defer res.Close()
	
	var status string 
	for res.Next(){
		err = res.Scan(&status)
		if err!=nil{
			return  "",err
		}
	}

	return status,nil
}