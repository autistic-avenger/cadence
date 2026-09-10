package db

import "database/sql"

type YoutubeDbInfo struct {
	Email     string `json:"email"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
	Creator   string `json:"creator"`
	JobID     string `json:"jobID"`
	VideoId	  string `json:"videoId"`
}

func AddVideoInfo(dbCursor *sql.DB, info YoutubeDbInfo) error {

	query := `
	INSERT INTO userVideos (email,jobid,title,thumbnail,url,creator,videoid)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	`
	row := dbCursor.QueryRow(query, info.Email, info.JobID, info.Title, info.Thumbnail, info.Url, info.Creator,info.VideoId)
	
	err := row.Err()
	if err!=nil{
		return err
	}
	return nil
}
