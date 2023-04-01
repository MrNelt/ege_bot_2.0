package database

import (
	"log"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/models"
	"gorm.io/gorm"
)

func FindUser(id uint) (*gorm.DB, models.User) {
	DB := GetDB()
	user := models.User{}
	result := DB.First(&user, id)
	return result, user
}

func CreateUser(id uint, name string, record uint) {
	DB := GetDB()
	user := models.User{Name: name, ID: id, Record: record, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Can't create user")
	}
}
