package database

import (
	"log"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/model"
	models "github.com/kappaprideonly/ege_bot_2.0/model"
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

func UpdateRecordUser(id, record uint) {
	DB := GetDB()
	user := model.User{ID: id}
	DB.Model(&user).Updates(model.User{UpdatedAt: time.Now(), Record: record})
}

func GetUsersOrderedByRecord() []model.User {
	DB := GetDB()
	users := []model.User{}
	DB.Limit(10).Order("record desc").Find(&users)
	return users
}

func CountOfUsers() int64 {
	DB := GetDB()
	var count int64
	DB.Model(model.User{}).Count(&count)
	return count
}
