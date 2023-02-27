package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetMysqlConnectionUrl() string {
	dbPort := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	mysqlConn := fmt.Sprintf("%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, dbPort, dbName)

	return mysqlConn
}

func ConnectToDB() {
	var err error
	dsn := GetMysqlConnectionUrl()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("failed to connect to database : " + err.Error())
	}
}