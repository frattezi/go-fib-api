package services

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func FibRedis(n int, client *redis.Client, ctx context.Context) int {
	keyName := strconv.Itoa(n)

	if n == 1 || n == 0 {
		err := client.Set(ctx, keyName, n, 0).Err()
		if err != nil {
			panic("Error persisting data")
		}
	}

	_, err := client.Get(ctx, keyName).Result()

	if err == redis.Nil {
		calculatedValue := FibRedis(n-1, client, ctx) + FibRedis(n-2, client, ctx)
		client.Set(ctx, keyName, calculatedValue, 0)
		return calculatedValue
	}

	value, err := client.Get(ctx, keyName).Result()

	if err != nil {
		panic(err)
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return result

}
