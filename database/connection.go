package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("USER_DB"),
		os.Getenv("PASS_DB"),
		os.Getenv("NAME_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("SSLMODE"))
	log.Printf("[postgres] %s", dsl)
	db, err := gorm.Open(postgres.Open(dsl), &gorm.Config{})
	if err != nil {
		log.Panic("Can't connect to db")
	}
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
	}
	return dbase
}
