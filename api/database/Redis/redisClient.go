package database

import (
	"context"
	"errors"
	"log"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Fatal(errors.New("REDIS_CONNECTION"))
	}

	return client
}
