package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/model"
)

func init() {
	if err := godotenv.Load("../config/.env"); err != nil {
		log.Panic("No .env file found")
	}
}

func main() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
	db.Delete(&model.User{}, "name LIKE ?", "Leeroy")
}
