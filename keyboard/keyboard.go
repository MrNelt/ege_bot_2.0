package keyboard

import tele "gopkg.in/telebot.v3"

var menuKeyboard *tele.ReplyMarkup
var trainingKeyboard *tele.ReplyMarkup

func CreateAllKeyboards() {
	menuKeyboard = &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := menuKeyboard.Text("1")
	btn2 := menuKeyboard.Text("2")
	btn3 := menuKeyboard.Text("3")
	btn4 := menuKeyboard.Text("4")
	menuKeyboard.Reply(
		menuKeyboard.Row(btn1, btn2),
		menuKeyboard.Row(btn3, btn4),
	)
}

func GetMenuKeyboard() *tele.ReplyMarkup {
	return menuKeyboard
}
