package middlewares

import (
	log "github.com/bearatol/lg"

	tele "gopkg.in/telebot.v3"
)

func Logger() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			log.Tracef("%s : %s", ctx.Sender().FirstName, ctx.Text())
			return next(ctx)
		}
	}
}
