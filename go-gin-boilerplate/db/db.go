package db

import (
	"github.com/vsouza/go-gin-boilerplate/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *dynamodb.DynamoDB

func Init() {
	c := config.GetConfig()

	DBMS := c.GetString("DB.dbms")
	HOST := c.GetString("DB.host")
	USER := c.GetString("DB.user")
	PASS := c.GetString("DB.pass")
	DBNAME := c.GetString("DB.dbname")
	SSLMODE := c.GetString("DB.sslmode")

	CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USER + " dbname=" + DBNAME + " password=" + PASS + " sslmode=" + SSLMODE

	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
}

func GetDB() *dynamodb.DynamoDB {
	return db
}
