package middlewares

import (
	"github.com/kappaprideonly/ege_bot_2.0/manager/session"
	tele "gopkg.in/telebot.v3"
)

func RedisSession() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			_, err := session.GetToken(uint(ctx.Sender().ID))
			if err != nil {
				session.CreateToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
			}
			return next(ctx)
		}
	}
}
