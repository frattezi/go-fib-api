package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis"
)

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result := rdb.Ping()
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf("{result: %s}", result.String()))
}
