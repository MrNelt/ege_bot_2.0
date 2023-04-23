package main

import (
	log "github.com/bearatol/lg"
	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/bot"
	"github.com/kappaprideonly/ege_bot_2.0/handlers"
	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
	"github.com/kappaprideonly/ege_bot_2.0/middlewares"
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

	bot.Use(middlewares.Logger())
	bot.Use(middlewares.OnlyPrivate())

	bot.Handle("/start", handlers.Auth)

	bot.Handle("/record", handlers.Record)
	bot.Handle("/leaderboard", handlers.LeaderBoard)
	bot.Handle("/stats", handlers.Stats)
	bot.Handle("/help", handlers.Help)

	bot.Handle("/menu", handlers.Menu)
	bot.Handle("/begin", handlers.Begin)
	bot.Handle(tele.OnText, handlers.ProcessTraining)

	adminOnly := bot.Group()
	adminOnly.Use(middlewares.OnlyAdmin())
	adminOnly.Handle("/admin", handlers.AdminTest)

	bot.Start()
}
