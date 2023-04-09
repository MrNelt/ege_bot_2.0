package main

import (
	"github.com/kappaprideonly/ege_bot_2.0/config"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/model"
)

func init() {
	config.Init()
}

func main() {
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
}
