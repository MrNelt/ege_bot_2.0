package handler

import (
	"fmt"

	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
	tele "gopkg.in/telebot.v3"
)

func Auth(ctx tele.Context) error {
	if storage.ExistUser(uint(ctx.Sender().ID)) {
		message := fmt.Sprintf("👍 <b>%s</b>, Вы уже зарегистрированы!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
		session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
		if err != nil {
			go sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
		} else {
			MenuSession(&session)
			go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
		}
		return ctx.Send(message, keyboard.GetMenuKeyboard())
	}
	go storage.AddUser(uint(ctx.Sender().ID), 0, ctx.Sender().FirstName)
	message := fmt.Sprintf("✅ <b>%s</b>, Вы успешно зарегистрированы!\nНачать тренировку - <b>/begin</b>", ctx.Sender().FirstName)
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		go sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	} else {
		MenuSession(&session)
		go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	}
	return ctx.Send(message+"\n"+TextHelp, keyboard.GetMenuKeyboard())
}
