package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
	tele "gopkg.in/telebot.v3"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Panic("No .env file found")
	}
}

func main() {
	key, exist := os.LookupEnv("KEY_BOT")
	log.Printf("%s\n", key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}
	database.Init()
	redisdb.Init()

	pref := tele.Settings{
		Token:  key,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Start()
}
