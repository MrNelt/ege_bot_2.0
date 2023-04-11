package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
	tele "gopkg.in/telebot.v3"
)

func Logger(logger ...*log.Logger) tele.MiddlewareFunc {
	var l *log.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(ctx tele.Context) error {
			data, _ := json.MarshalIndent(ctx.Update(), "", "  ")
			l.Println(string(data))
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
			_, err := redisdb.ReceiveToken(uint(ctx.Sender().ID))
			if err != nil {
				redisdb.NewToken(uint(ctx.Sender().ID), ctx.Sender().FirstName)
			}
			return next(ctx)
		}
	}
}
