package session

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kappaprideonly/ege_bot_2.0/manager/model"
	"github.com/kappaprideonly/ege_bot_2.0/manager/storage"
	"github.com/redis/go-redis/v9"
)

type RedisSessionDB struct {
	client      *redis.Client
	ctx         *context.Context
	timeSession int
}

func NewRedisSessionDB(host, pass string, timeSession int) *RedisSessionDB {
	log.Printf("[redis] %s, %s", host, pass)
	db := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		PoolSize: 1000,
		DB:       0,
	})
	ctx := context.Background()
	_, err := db.Ping(ctx).Result()
	if err != nil {
		log.Panic("Can't connect to Redis")
	}
	return &RedisSessionDB{db, &ctx, timeSession}
}

func (r *RedisSessionDB) getToken(id uint) (model.Token, error) {
	conn := r.client
	js, err := conn.Get(*r.ctx, fmt.Sprint(id)).Result()
	if err != nil {
		return model.Token{}, err
	}
	user := model.Token{}
	json.Unmarshal([]byte(js), &user)
	return user, nil
}

func (r *RedisSessionDB) updateToken(id uint, user model.Token) error {
	js, err := json.Marshal(user)
	if err != nil {
		log.Printf("[redis] Can't create json")
	}
	conn := r.client
	err = conn.Set(*r.ctx, fmt.Sprint(id), js, time.Duration(r.timeSession)*time.Minute).Err()
	return err
}

func (r *RedisSessionDB) createToken(id uint, name string) model.Token {
	user := storage.FindUser(id, name)
	token := model.Token{}
	token.Answer = ""
	token.Condition = "new"
	token.CurrentScore = 0
	token.Record = user.Record
	r.updateToken(id, token)
	return token
}
