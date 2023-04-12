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

func FindUser(id uint, name string) models.User {
	DB := GetDB()
	user := models.User{}
	if result := DB.First(&user, id); result.Error != nil {
		CreateUser(id, name, 0)
		DB.First(&user, id)
	}
	return user
}

func CreateUser(id uint, name string, record uint) {
	DB := GetDB()
	user := models.User{Name: name, ID: id, Record: record, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Panic("Can't create user")
	}
}

func UpdateRecordUser(id, record uint, name string) {
	DB := GetDB()
	user := model.User{ID: id}
	if result := DB.Model(&user).Updates(model.User{UpdatedAt: time.Now(), Record: record}); result.RowsAffected == 0 {
		CreateUser(id, name, record)
	}
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
