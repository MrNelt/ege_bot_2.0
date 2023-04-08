package task

import "github.com/kappaprideonly/ege_bot_2.0/model"

func GetTask() model.Task {
	task := model.Task{}
	task.Answer = "правильный"
	task.Wrong[0] = "неправ1"
	task.Wrong[1] = "неправ2"
	task.Wrong[2] = "неправ3"
	return task
}
