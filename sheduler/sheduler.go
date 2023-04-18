package sheduler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
)

var leaderboard string
var count int64

func Init() {
	updateMinutes, err := strconv.Atoi(os.Getenv("UPTIME_MIN"))
	if err != nil {
		log.Panic("Can't parse to int UPTIME_MIN")
	}

	sheduler := gocron.NewScheduler(time.Now().Location())
	sheduler.Every(updateMinutes).Minute().Do(updateLeaderboard)
	sheduler.Every(updateMinutes).Minute().Do(updateCount)
	sheduler.StartAsync()
}

func updateLeaderboard() {
	log.Print("[sheduler] update leaderboard")
	users := storage.GetUsersOrderedByRecord()
	leaderboard = "🏆 Таблица лидеров:\n\n"
	for i, v := range users {
		if i == 0 {
			leaderboard += fmt.Sprintf("🥇 %s - [%d]\n", v.Name, v.Record)
		} else if i == 1 {
			leaderboard += fmt.Sprintf("🥈 %s - [%d]\n", v.Name, v.Record)
		} else if i == 2 {
			leaderboard += fmt.Sprintf("🥉 %s - [%d]\n", v.Name, v.Record)
		} else {
			leaderboard += fmt.Sprintf(" %d)   %s - [%d]\n", i+1, v.Name, v.Record)
		}
	}
}

func GetLeaderboard() string {
	return leaderboard
}

func updateCount() {
	log.Print("[sheduler] update count")
	count = storage.CountOfUsers()
}

func GetCount() int64 {
	return count
}
