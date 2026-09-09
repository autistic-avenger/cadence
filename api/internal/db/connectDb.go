package db

import (
	"database/sql"
	"os"
)

func ConnectDB() (*sql.DB,error) {

	dbCursor, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		return nil,err
	}
	
	err = dbCursor.Ping()
	if err != nil {
		return nil,err
	}
	
	return dbCursor,nil
}