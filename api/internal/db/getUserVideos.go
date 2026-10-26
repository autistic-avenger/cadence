package db

import (
	"database/sql"
)

func DbGetVideos(dbCursor *sql.DB, email string)([]YoutubeDbInfo,error){
	var vids []YoutubeDbInfo
	query := `SELECT email, jobid, title, thumbnail, url, creator, videoid, jobstatus FROM userVideos WHERE email = $1 ORDER BY created_at DESC`

	rows,err := dbCursor.Query(query,email)
	if err!=nil{
		return vids,err
	}
	defer rows.Close()

	for rows.Next(){
		var vid YoutubeDbInfo
		err := rows.Scan(
			&vid.Email,
			&vid.JobID,
			&vid.Title,
			&vid.Thumbnail,
			&vid.Url,
			&vid.Creator,
			&vid.VideoId,
			&vid.JobStatus,
		)
		if err!=nil{
			return nil,err
		}
		vids = append(vids,vid)
	}

	return vids,nil
}