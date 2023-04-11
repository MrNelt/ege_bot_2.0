package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/kappaprideonly/ege_bot_2.0/database"
	"github.com/kappaprideonly/ege_bot_2.0/model"
	"github.com/kappaprideonly/ege_bot_2.0/redisdb"
)

func init() {
	if err := godotenv.Load("../config/.env"); err != nil {
		log.Panic("No .env file found")
	}
}

func BenchmarkRedis(b *testing.B) {
	for j := 0; j < b.N; j++ {
		id := []string{}
		for i := 0; i < 50; i++ {
			key := rand.Uint64()
			tmpKey := fmt.Sprintf("%d", key)
			redisdb.UpdateToken(uint(key), model.Token{})
			id = append(id, tmpKey)
		}
		for i := 0; i < len(id); i++ {
			key, _ := strconv.Atoi(id[i])
			redisdb.ReceiveToken(uint(key))
		}
		client := redisdb.GetClient()
		client.Del(redisdb.GetCtx(), id...)
	}
}

func BenchmarkPostgres(b *testing.B) {

	for j := 0; j < b.N; j++ {
		id := []uint32{}
		for i := 0; i < 50; i++ {
			key := rand.Uint32()
			database.CreateUser(uint(key), "", 0)
			id = append(id, key)
		}
		for i := 0; i < len(id); i++ {
			database.FindUser(uint(id[i]), "")
		}
		db := database.GetDB()
		db.Delete(&model.User{}, "Name LIKE ?", "")
	}
}
