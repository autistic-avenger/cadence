package db

import "database/sql"

func CreateTable(dbCursor *sql.DB) error {
	createTable := `
		CREATE TABLE IF NOT EXISTS userVideos (
			email VARCHAR(50),
			jobid VARCHAR(36),
			title VARCHAR(100),
			thumbnail VARCHAR(70),
			url VARCHAR(70),
			creator VARCHAR(60),
			videoid VARCHAR(20)
		)`

	_, err := dbCursor.Exec(createTable)
	if err !=nil{
		return err
	}
	return nil
}