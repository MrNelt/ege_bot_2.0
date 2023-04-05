package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/models"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
)

func main() {
	key, exist := os.LookupEnv("KEY_BOT")
	log.Printf("%s\n", key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}
	database.Init()
	redisdb.Init()
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			log.Printf("%d", update.Message.From.ID)
			id := update.Message.From.ID
			name := update.Message.From.FirstName
			if update.Message.Text == "/start" {
				result, _ := database.FindUser(uint(id))
				if result.Error != nil {
					log.Printf("Can't find user with id=%d", id)
					database.CreateUser(uint(id), name, 0)
				} else {
					log.Printf("User find!")
					bot.Send(tgbotapi.NewMessage(id, "Вы авторизованы!"))
				}
			} else {
				token, err := redisdb.ReceiveToken(uint(id))
				if err != nil {
					user := models.Token{CurrentScore: 0, Answer: "test", Condition: "test", Record: 0}
					redisdb.UpdateToken(uint(id), user)
					bot.Send(tgbotapi.NewMessage(id, "Сессия создана!"))

				} else {
					message := fmt.Sprintf("%v", token)
					log.Printf(message)
					bot.Send(tgbotapi.NewMessage(id, message))
				}
			}
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID
			// bot.Send(msg)
		}
	}
}
