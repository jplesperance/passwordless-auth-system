package rdb

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisAddress == "" || redisPassword == "" {
		log.Fatal("Missing required redis environment variables")
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
	})
}
