package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/frattezi/go-fib-api/services"
)

func FibService(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("n"))

	if err != nil {
		panic(err.Error())
	}

	result := services.Fib(n)
	fmt.Fprintf(w, "Fibonacci %v = %v\n", n, result)
}
