package middlewares

import (
	"fmt"
	"os"

	tele "gopkg.in/telebot.v3"
)

func OnlyAdmin() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			if fmt.Sprintf("%d", ctx.Sender().ID) != os.Getenv("ADMIN_ID") {
				return nil
			}
			return next(ctx)
		}
	}
}
