package db

import "database/sql"

type YoutubeDbInfo struct {
	Email     string
	Title     string
	Thumbnail string
	Url       string
	Creator   string
	JobID     string
}

func AddVideoInfo(dbCursor *sql.DB, info YoutubeDbInfo) error {

	query := `
	INSERT INTO userVideos (email,jobid,title,thumbnail,url,creator)
	VALUES ($1,$2,$3,$4,$5,$6)
	`
	row := dbCursor.QueryRow(query, info.Email, info.JobID, info.Title, info.Thumbnail, info.Url, info.Creator)
	
	err := row.Err()
	if err!=nil{
		return err
	}
	return nil
}
