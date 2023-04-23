package middlewares

import (
	tele "gopkg.in/telebot.v3"
)

func OnlyPrivate() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if ctx.Message().Chat.Type != "private" {
				return nil
			}
			return next(ctx)
		}
	}
}
