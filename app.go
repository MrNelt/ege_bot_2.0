package main

import (
	log "github.com/bearatol/lg"
	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/bot"
	"github.com/kappaprideonly/ege_bot_2.0/handler"
	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
	"github.com/kappaprideonly/ege_bot_2.0/middleware"
	"github.com/kappaprideonly/ege_bot_2.0/sheduler"
	"github.com/kappaprideonly/ege_bot_2.0/task"
	tele "gopkg.in/telebot.v3"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
	task.Init()
	storage.Init()
	keyboard.Init()
	sessionDB.Init()
	sheduler.Init()
}

func main() {

	bot := bot.NewBot()

	bot.Use(middleware.Logger())
	bot.Use(middleware.OnlyPrivate())

	bot.Handle("/start", handler.Auth)

	bot.Handle("/record", handler.Record)
	bot.Handle("/leaderboard", handler.LeaderBoard)
	bot.Handle("/stats", handler.Stats)
	bot.Handle("/help", handler.Help)

	bot.Handle("/menu", handler.Menu)
	bot.Handle("/begin", handler.Begin)
	bot.Handle(tele.OnText, handler.ProcessTraining)

	adminOnly := bot.Group()
	adminOnly.Use(middleware.OnlyAdmin())
	adminOnly.Handle("/admin", handler.AdminTest)

	bot.Start()
}
