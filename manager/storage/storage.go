package storage

import (
	"fmt"
	"os"

	log "github.com/bearatol/lg"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
)

type Storage interface {
	existUser(id uint) bool
	addUser(id, record uint, name string)
	findUser(id uint, name string) model.User
	countOfUsers() int64
	updateRecordUser(id, record uint, name string)
	getUsersOrderedByRecord() []model.User
}

var storage Storage

func Init() {
	dsl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"))
	log.Debug("[postgres] %s", dsl)
	storage = NewPgStorage(dsl)
}

func ExistUser(id uint) bool {
	return storage.existUser(id)
}

func FindUser(id uint, name string) model.User {
	return storage.findUser(id, name)
}

func AddUser(id, record uint, name string) {
	storage.addUser(id, record, name)
}

func UpdateRecordUser(id, record uint, name string) {
	storage.updateRecordUser(id, record, name)
}

func GetUsersOrderedByRecord() []model.User {
	return storage.getUsersOrderedByRecord()
}

func CountOfUsers() int64 {
	return storage.countOfUsers()
}
