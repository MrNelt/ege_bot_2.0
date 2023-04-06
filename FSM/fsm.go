package FSM

import (
	"github.com/kappaprideonly/ege_bot_2.0/models"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
)

func FSM(id uint, message string) string {
	_, err := redisdb.ReceiveToken(id)
	if err != nil {

		redisdb.UpdateToken(id, models.Token{})
	}
	return "tests"
}
