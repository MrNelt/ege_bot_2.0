package redisdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/model"
)

var minutes int

func GetMinutes() int {
	if minutes != 0 {
		return minutes
	}
	minutes, errParse := strconv.Atoi(os.Getenv("SESSION_TIME_MIN"))
	if errParse != nil {
		log.Printf("[redis] Can't parse SESSION_TIME_MIN to int")
	}
	return minutes
}

func ReceiveToken(id uint) (model.Token, error) {
	conn := GetClient()
	js, err := conn.Get(GetCtx(), fmt.Sprint(id)).Result()
	if err != nil {
		return model.Token{}, err
	}
	user := model.Token{}
	json.Unmarshal([]byte(js), &user)
	return user, nil

}

func UpdateToken(id uint, user model.Token) error {
	js, err := json.Marshal(user)
	if err != nil {
		log.Printf("[redis] Can't create json")
	}
	conn := GetClient()
	err = conn.Set(GetCtx(), fmt.Sprint(id), js, time.Duration(GetMinutes())*time.Minute).Err()
	return err
}

func NewToken(id uint) model.Token {
	user := database.FindUser(id)
	token := model.Token{}
	token.Answer = ""
	token.Condition = "new"
	token.CurrentScore = 0
	token.Record = user.Record
	UpdateToken(id, token)
	return token
}
