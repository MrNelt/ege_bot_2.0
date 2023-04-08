package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/middlewares"
	"github.com/kappaprideonly/ege_bot_2.0/models"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
	"github.com/kappaprideonly/ege_bot_2.0/tasks"
	tele "gopkg.in/telebot.v3"
)

func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Panic("No .env file found")
	}
	database.Init()
	redisdb.Init()
}

func DefaultSession(session *models.Token) {
	session.Answer = ""
	session.Condition = "menu"
	session.CurrentScore = 0
}

func BeginTrainingSession(session *models.Token, task models.Task) {
	session.Answer = task.Answer
	session.Condition = "training"
	session.CurrentScore = 0
}

func main() {
	key, exist := os.LookupEnv("KEY_BOT")
	log.Printf("[key] %s\n", key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}

	pref := tele.Settings{
		Token:  key,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Use(middlewares.Logger())
	bot.Use(middlewares.OnlyPrivate())

	bot.Handle("/start", func(ctx tele.Context) error {
		if database.ExistUser(uint(ctx.Sender().ID)) {
			message := fmt.Sprintf("🖐🏾 %s, Вы уже зарегистрированы!\nНачать тренировку - /begin", ctx.Sender().FirstName)
			return ctx.Send(message)
		}
		go database.CreateUser(uint(ctx.Sender().ID), ctx.Sender().FirstName)
		message := fmt.Sprintf("🖐🏾 %s, Вы успешно зарегистрированы!\nНачать тренировку - /begin", ctx.Sender().FirstName)
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			go redisdb.NewToken(uint(ctx.Sender().ID))
		} else {
			DefaultSession(&session)
			go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		}
		return ctx.Send(message)
	})

	authOnly := bot.Group()
	authOnly.Use(middlewares.RedisSession())

	authOnly.Handle("/record", func(ctx tele.Context) error {
		session, _ := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		record := fmt.Sprintf("%d", session.Record)
		return ctx.Send(record)
	})

	authOnly.Handle("/begin", func(ctx tele.Context) error {
		session, _ := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		task := tasks.GetTask()
		BeginTrainingSession(&session, task)
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		message := fmt.Sprintf("%v", task)
		return ctx.Send(message)
	})

	adminOnly := bot.Group()
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.Handle("/admin", func(ctx tele.Context) error {
		return ctx.Send("test")
	})

	bot.Start()
}
