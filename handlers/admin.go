package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func AdminTest(ctx tele.Context) error {
	return ctx.Send("test")
}
