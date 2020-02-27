package db

import (
	"fmt"
	"log"
	"os"
	//"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// LogCheckin will increment the amount of times that someone has checked in
func LogCheckin() {
	// TODO: Maybe log the user who had logged in? idk
	db := openDb()
	defer db.Close()
}

// GetCheckin will return the amount of times checkin has been run
func GetCheckins() {
	// TODO: Get checkins for a specific user
	db := openDb()
	defer db.Close()
}

func openDb () *gorm.DB {
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