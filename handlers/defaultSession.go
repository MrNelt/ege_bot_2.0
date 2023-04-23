package handlers

import "github.com/kappaprideonly/ege_bot_2.0/manager/model"

func MenuSession(session *model.Token) {
	session.Answer = ""
	session.Condition = "menu"
	session.CurrentScore = 0
}

func BeginTrainingSession(session *model.Token, task model.Task) {
	session.Answer = task.Answer
	session.Condition = "training"
	session.CurrentScore = 0
}
