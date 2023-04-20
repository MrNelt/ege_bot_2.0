package handler

import (
	"fmt"

	"github.com/kappaprideonly/ege_bot_2.0/keyboard"
	sessionDB "github.com/kappaprideonly/ege_bot_2.0/manager/session"
	"github.com/kappaprideonly/ege_bot_2.0/sheduler"
	tele "gopkg.in/telebot.v3"
)

func Record(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	record := fmt.Sprintf("💪 Ваш рекорд: %d", session.Record)
	if session.Condition == "training" {
		return ctx.Send(record, keyboard.GetTrainingKeyboard())
	}
	return ctx.Send(record, keyboard.GetMenuKeyboard())
}

func LeaderBoard(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	leaderboard := sheduler.GetLeaderboard()
	if session.Condition == "training" {
		return ctx.Send(leaderboard, keyboard.GetTrainingKeyboard())
	}
	return ctx.Send(leaderboard, keyboard.GetMenuKeyboard())
}

func Stats(ctx tele.Context) error {
	session, err := sessionDB.GetToken(uint(ctx.Sender().ID))
	if err != nil {
		session = sessionDB.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
	}
	go sessionDB.UpdateToken(uint(ctx.Sender().ID), session)
	count := sheduler.GetCount()
	message := fmt.Sprintf("📊 Количество пользователей: %d", count)
	if session.Condition == "training" {
		return ctx.Send(message, keyboard.GetTrainingKeyboard())
	}
	return ctx.Send(message, keyboard.GetMenuKeyboard())
}

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
