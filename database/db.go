package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_DATABASE = "yuanaturego"
	DB_USERNAME = "root"
	DB_PASSWORD = ""
)

var Db *gorm.DB

func init() {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB is connected")
}
