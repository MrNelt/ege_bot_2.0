package handlers

import (
	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/task"
	tele "gopkg.in/telebot.v3"
)


func Begin(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	task, message := task.GetTask()
	BeginTrainingSession(&session, task)
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	return ctx.Send(message, keyboard.GetTrainingKeyboard())
}