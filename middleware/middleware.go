package middleware

import (
	"fmt"
	"os"

	log "github.com/bearatol/lg"

	"github.com/kappaprideonly/ege_bot_2.0/manager/session"
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
