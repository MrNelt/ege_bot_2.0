package bot

import (
	"os"
	"time"

	log "github.com/bearatol/lg"
	tele "gopkg.in/telebot.v3"
)

func NewBot() *tele.Bot {

	key, exist := os.LookupEnv("KEY_BOT")
	log.Debugf("[key] %s\n", key)
	if exist == false {
		log.Panic("Key doesn't exist")
	}
	pref := tele.Settings{
		Token:     key,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Panic(err)
	}
	return bot
}
