package services

import (
	"strconv"

	"github.com/go-redis/redis"
)

func FibRedis(n int, client *redis.Client) int {
	keyName := strconv.Itoa(n)

	if n == 1 || n == 0 {
		client.Set(keyName, n, 0)
	}

	persistedValue, err := client.Get(keyName).Result()
	if err != nil {
		panic(err)
	}

	if err == redis.Nil {
		result := FibRedis(n-1, client) + FibRedis(n-2, client)
		client.Set(keyName, result, 0)
	}
	result, err := strconv.Atoi(persistedValue)

	if err != nil {
		panic(err)
	}

	return result
}
