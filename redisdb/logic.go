package redisdb

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/models"
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

func ReceiveToken(id uint) (models.Token, error) {
	client := GetClient()
	js, err := client.Get(GetCtx(), string(rune(id))).Result()
	if err != nil {
		return models.Token{}, err
	}
	user := models.Token{}
	json.Unmarshal([]byte(js), &user)
	return user, nil
}

func UpdateToken(id uint, user models.Token) error {
	js, err := json.Marshal(user)
	if err != nil {
		log.Printf("[redis] Can't create json")
	}
	client := GetClient()

	err = client.Set(GetCtx(), string(rune(id)), js, time.Duration(GetMinutes())*time.Minute).Err()
	return err
}
