package db

import (
	"database/sql"
)

func DbGetVideos(dbCursor *sql.DB, email string)([]YoutubeDbInfo,error){
	var vids []YoutubeDbInfo
	query := `SELECT * FROM userVideos WHERE email = $1`

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
		)
		if err!=nil{
			return nil,err
		}
		vids = append(vids,vid)
	}

	return vids,nil
}