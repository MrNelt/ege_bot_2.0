package main

import (
	"fmt"
	"os"

	log "github.com/bearatol/lg"

	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load("../config/.env"); err != nil {
		log.Panic("No .env file found")
	}
}

func main() {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("HOST_DB"),
		os.Getenv("USER_DB"),
		os.Getenv("PASS_DB"),
		os.Getenv("NAME_DB"),
		os.Getenv("PORT_DB"),
		os.Getenv("SSLMODE"))
	log.Debug("[postgres] %s", dsl)
	db, err := gorm.Open(postgres.Open(dsl), &gorm.Config{})
	if err != nil {
		log.Panic("Can't connect to db")
	}
	db.AutoMigrate(&model.User{})
	db.Delete(&model.User{}, "name LIKE ?", "Leeroy")
}
