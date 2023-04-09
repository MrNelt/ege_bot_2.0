package keyboard

import tele "gopkg.in/telebot.v3"

var menuKeyboard *tele.ReplyMarkup
var trainingKeyboard *tele.ReplyMarkup

func CreateAllKeyboards() {
	trainingKeyboard = &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := trainingKeyboard.Text("1")
	btn2 := trainingKeyboard.Text("2")
	btn3 := trainingKeyboard.Text("3")
	btn4 := trainingKeyboard.Text("4")
	trainingKeyboard.Reply(
		trainingKeyboard.Row(btn1, btn2),
		trainingKeyboard.Row(btn3, btn4),
	)
	menuKeyboard = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnBegin := menuKeyboard.Text("/begin")
	btnRecord := menuKeyboard.Text("/record")
	btnLeader := menuKeyboard.Text("/leaderboard")
	btnStats := menuKeyboard.Text("/stats")
	menuKeyboard.Reply(
		menuKeyboard.Row(btnBegin, btnStats),
		menuKeyboard.Row(btnRecord, btnLeader),
	)
}

func GetMenuKeyboard() *tele.ReplyMarkup {
	return menuKeyboard
}

func GetTrainingKeyboard() *tele.ReplyMarkup {
	return trainingKeyboard
}
