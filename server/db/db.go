package db

import (
	"fmt"
	"log"
	"os"
	//"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() *gorm.DB {
	sqlHost := os.Getenv("DB_HOST")
	sqlUser := os.Getenv("DB_USER")
	sqlPassword := os.Getenv("DB_PASSWORD")
	sqlDb := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", sqlUser, sqlPassword, sqlHost, sqlDb)

	db, err := gorm.Open("mysql", connString)

	if err != nil {
		log.Fatal(err)
	}

	return db
}