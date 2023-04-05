package redisdb

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client *redis.Client

func Init() *redis.Client {
	log.Printf("[redis] %s, %s", os.Getenv("HOST_REDIS"), os.Getenv("PASS_REDIS"))
	db := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("HOST_REDIS"),
		Password: os.Getenv("PASS_REDIS"),
		DB:       0,
	})
	_, err := db.Ping(ctx).Result()
	if err != nil {
		log.Panic("Can't connect to Redis")
	}
	return db
}

func GetClient() *redis.Client {
	if client == nil {
		client = Init()
	}
	return client
}

func GetCtx() context.Context {
	return ctx
}
