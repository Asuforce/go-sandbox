package model

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := "gwa"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
