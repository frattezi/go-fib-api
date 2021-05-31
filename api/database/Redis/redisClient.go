package database

import (
	"errors"
	"log"

	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(errors.New("REDIS_CONNECTION"))
	}

	return client
}
