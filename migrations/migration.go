package main

import (
	"fmt"
	"os"

	log "github.com/bearatol/lg"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSGTRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("POSTGRES_SSLMODE"))
	log.Debug("[postgres] %s", dsl)
	db, err := gorm.Open(postgres.Open(dsl), &gorm.Config{})
	if err != nil {
		log.Panic("Can't connect to db")
	}
	db.AutoMigrate(&model.User{})
	db.Delete(&model.User{}, "name LIKE ?", "Leeroy")
}
