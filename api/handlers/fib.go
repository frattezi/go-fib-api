package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	database "github.com/frattezi/go-fib-api/database/Redis"
	"github.com/frattezi/go-fib-api/services"
)

var ctx = context.Background()

func FibHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	if err != nil || n < 0 {
		http.Error(w, "INVALID_ENTRY", http.StatusUnprocessableEntity)
		return
	}

	result := services.Fib(n)

	// Change to JSON object for response
	io.WriteString(w, fmt.Sprintf("{result: %d}", result))
}

func FibRedisHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	if err != nil || n < 0 {
		http.Error(w, "INVALID_ENTRY", http.StatusUnprocessableEntity)
		return
	}
	client := database.GetRedisClient(ctx)

	//lookup := make(map[int]int)
	result := services.FibRedis(n, client, ctx)

	// Change to JSON object for response
	io.WriteString(w, fmt.Sprintf("{result: %d}", result))
}
