package handlers

import (
	"fmt"

	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	tele "gopkg.in/telebot.v3"
)

func Menu(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	MenuSession(&session)
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	message := fmt.Sprintf("🪖 <b>%s</b>, Вы в главном меню!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
	return ctx.Send(message, keyboard.GetMenuKeyboard())
}
