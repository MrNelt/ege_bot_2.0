package handlers

import (
	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	tele "gopkg.in/telebot.v3"
)

func Help(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	message := TextHelp
	if session.Condition == "training" {
		return ctx.Send(message, keyboard.GetTrainingKeyboard())
	}
	return ctx.Send(message, keyboard.GetMenuKeyboard())
}
