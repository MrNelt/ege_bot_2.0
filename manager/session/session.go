package session

import (
	"log"
	"os"
	"strconv"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
)

type SessionDB interface {
	createToken(id uint, name string) model.Token
	getToken(id uint) (model.Token, error)
	updateToken(id uint, user model.Token) error
}

var sessionDB SessionDB

func Init() {
	host := os.Getenv("HOST_REDIS")
	pass := os.Getenv("PASS_REDIS")
	minutes, errParse := strconv.Atoi(os.Getenv("SESSION_TIME_MIN"))
	if errParse != nil {
		log.Panic("[redis] Can't parse SESSION_TIME_MIN to int")
	}
	sessionDB = NewRedisSessionDB(host, pass, minutes)
}

func CreateToken(id uint, name string) model.Token {
	return sessionDB.createToken(id, name)
}

func GetToken(id uint) (model.Token, error) {
	return sessionDB.getToken(id)
}

func UpdateToken(id uint, user model.Token) error {
	return sessionDB.updateToken(id, user)
}
