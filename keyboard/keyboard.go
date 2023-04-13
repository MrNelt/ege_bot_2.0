package keyboard

import tele "gopkg.in/telebot.v3"

type keyboard interface {
	init()
	get() *tele.ReplyMarkup
}

type TrainingKeyboard struct {
	keyboard *tele.ReplyMarkup
}

func (k TrainingKeyboard) get() *tele.ReplyMarkup {
	return k.keyboard
}

func (k *TrainingKeyboard) init() {
	k.keyboard = &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := k.keyboard.Text("1")
	btn2 := k.keyboard.Text("2")
	btn3 := k.keyboard.Text("3")
	btn4 := k.keyboard.Text("4")
	k.keyboard.Reply(
		k.keyboard.Row(btn1, btn2),
		k.keyboard.Row(btn3, btn4),
	)
}

type MenuKeyboard struct {
	keyboard *tele.ReplyMarkup
}

func (k MenuKeyboard) get() *tele.ReplyMarkup {
	return k.keyboard
}

func (k *MenuKeyboard) init() {
	k.keyboard = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnBegin := k.keyboard.Text("/begin")
	btnRecord := k.keyboard.Text("/record")
	btnLeader := k.keyboard.Text("/leaderboard")
	btnStats := k.keyboard.Text("/stats")
	k.keyboard.Reply(
		k.keyboard.Row(btnBegin, btnStats),
		k.keyboard.Row(btnRecord, btnLeader),
	)
}

var menuKeyboard MenuKeyboard
var trainingKeyboard TrainingKeyboard

func Init() {
	menuKeyboard.init()
	trainingKeyboard.init()
}

func GetMenuKeyboard() *tele.ReplyMarkup {
	return menuKeyboard.get()
}

func GetTrainingKeyboard() *tele.ReplyMarkup {
	return trainingKeyboard.get()
}
