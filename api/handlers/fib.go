package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/frattezi/go-fib-api/services"
)

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
