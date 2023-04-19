package config

import (
	log "github.com/bearatol/lg"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Panic("No .env file found")
	}
}
