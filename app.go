package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/config"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	"github.com/kappaprideonly/ege_bot_2.0/middleware"
	"github.com/kappaprideonly/ege_bot_2.0/model"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
	"github.com/kappaprideonly/ege_bot_2.0/task"
	tele "gopkg.in/telebot.v3"
)

func init() {
	task.Init()
	config.Init()
	database.Init()
	keyboard.Init()
	redisdb.Init()
}

func MenuSession(session *model.Token) {
	session.Answer = ""
	session.Condition = "menu"
	session.CurrentScore = 0
}

func BeginTrainingSession(session *model.Token, task model.Task) {
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
		Token:     key,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Use(middleware.Logger())
	bot.Use(middleware.OnlyPrivate())

	bot.Handle("/start", func(ctx tele.Context) error {
		if database.ExistUser(uint(ctx.Sender().ID)) {
			message := fmt.Sprintf("🖐🏾 <b>%s</b>, Вы уже зарегистрированы!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
			session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
			if err != nil {
				go redisdb.NewToken(uint(ctx.Sender().ID))
			} else {
				MenuSession(&session)
				go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
			}
			return ctx.Send(message, keyboard.GetMenuKeyboard())
		}
		go database.CreateUser(uint(ctx.Sender().ID), ctx.Sender().FirstName)
		message := fmt.Sprintf("🖐🏾 <b>%s</b>, Вы успешно зарегистрированы!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			go redisdb.NewToken(uint(ctx.Sender().ID))
		} else {
			MenuSession(&session)
			go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		}
		return ctx.Send(message, keyboard.GetMenuKeyboard())
	})

	bot.Handle("/record", func(ctx tele.Context) error {
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			session = redisdb.NewToken(uint(ctx.Sender().ID))
		}
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		record := fmt.Sprintf("💪 Ваш рекорд: %d", session.Record)
		if session.Condition == "training" {
			return ctx.Send(record, keyboard.GetTrainingKeyboard())
		}
		return ctx.Send(record, keyboard.GetMenuKeyboard())
	})

	bot.Handle("/menu", func(ctx tele.Context) error {
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			session = redisdb.NewToken(uint(ctx.Sender().ID))
		}
		MenuSession(&session)
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		message := fmt.Sprintf("🪖 <b>%s</b>, Вы в главном меню!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
		return ctx.Send(message, keyboard.GetMenuKeyboard())
	})

	bot.Handle("/begin", func(ctx tele.Context) error {
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			session = redisdb.NewToken(uint(ctx.Sender().ID))
		}
		task, message := task.GetTask()
		BeginTrainingSession(&session, task)
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		return ctx.Send(message, keyboard.GetTrainingKeyboard())
	})

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		session, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
		if err != nil {
			session = redisdb.NewToken(uint(ctx.Sender().ID))
		}
		if session.Condition == "new" {
			return ctx.Send(fmt.Sprintf("⌛️ <b>%s</b>, Ваша сессия была окончена!\nНачать новую тренировку - <b>/begin</b>", ctx.Sender().FirstName), keyboard.GetMenuKeyboard())
		} else if session.Condition == "training" {
			if ctx.Text() != session.Answer {
				message := fmt.Sprintf("❌ Неверно!\nВаш score: <b>[%d]</b>\nНачать тренировку - <b>/begin</b>", session.CurrentScore)
				MenuSession(&session)
				go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
				return ctx.Send(message, keyboard.GetMenuKeyboard())
			}
			session.CurrentScore += 1
			if session.CurrentScore > session.Record {
				session.Record++
				go database.UpdateRecordUser(uint(ctx.Sender().ID), session.Record)
			}
			task, question := task.GetTask()
			session.Answer = task.Answer
			message := fmt.Sprintf("✅ Верно!\nВаш score: <b>[%d]</b>\nСледующее задание:\n%s", session.CurrentScore, question)
			go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
			return ctx.Send(message, keyboard.GetTrainingKeyboard())
		}
		go redisdb.UpdateToken(uint(ctx.Sender().ID), session)
		return ctx.Send(fmt.Sprintf("😤 <b>%s</b>, Я Вас не понимаю!", ctx.Sender().FirstName), keyboard.GetMenuKeyboard())
	})

	adminOnly := bot.Group()
	adminOnly.Use(middleware.OnlyAdmin())
	adminOnly.Handle("/admin", func(ctx tele.Context) error {
		return ctx.Send("test")
	})

	bot.Start()
}
