package handlers

import (
	"fmt"

	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
	"github.com/kappaprideonly/ege_bot_2.0/task"
	tele "gopkg.in/telebot.v3"
)



func ProcessTraining(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	if session.Condition == "new" {
		return ctx.Send(fmt.Sprintf("⌛️ <b>%s</b>, Ваша сессия была окончена!\nНачать новую тренировку - <b>/begin</b>", ctx.Sender().FirstName), keyboard.GetMenuKeyboard())
	} else if session.Condition == "training" {
		if ctx.Text() != session.Answer {
			message := fmt.Sprintf("❌ Неверно!\nВаш score: <b>[%d]</b>\nНачать тренировку - <b>/begin</b>", session.CurrentScore)
			MenuSession(&session)
			go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
			return ctx.Send(message, keyboard.GetMenuKeyboard())
		}
		session.CurrentScore += 1
		if session.CurrentScore > session.Record {
			session.Record++
			go storage.UpdateRecordUser(uint(ctx.Sender().ID), session.Record, ctx.Sender().FirstName)
		}
		task, question := task.GetTask()
		session.Answer = task.Answer
		message := fmt.Sprintf("✅ Верно!\nВаш score: <b>[%d]</b>\nСледующее задание:\n%s", session.CurrentScore, question)
		go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
		return ctx.Send(message, keyboard.GetTrainingKeyboard())
	}
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	return ctx.Send(fmt.Sprintf("😤 <b>%s</b>, Я Вас не понимаю!", ctx.Sender().FirstName), keyboard.GetMenuKeyboard())
}
