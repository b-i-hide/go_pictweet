package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitPictweetDB() *sql.DB {
	db, err := sql.Open("mysql", "root@/go_pictweet_development?parseTime=true")
	if err != nil {
		log.Fatalf("cannot open db, because %s", err)
		return nil
	}
	return  db
}
