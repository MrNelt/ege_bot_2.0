package database

import (
	"log"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/models"
)

func ExistUser(id uint) bool {
	DB := GetDB()
	user := models.User{}
	result := DB.First(&user, id)
	return result.Error == nil
}

func FindUser(id uint) models.User {
	DB := GetDB()
	user := models.User{}
	DB.First(&user, id)
	return user
}

func CreateUser(id uint, name string) {
	DB := GetDB()
	user := models.User{Name: name, ID: id, Record: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Panic("Can't create user")
	}
}
