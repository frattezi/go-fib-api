package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frattezi/go-fib-api/handlers"
	"github.com/frattezi/go-fib-api/middlewares"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
)

func handleRequests(port string) {
	// Route setup - fib
	fibHandler := http.HandlerFunc(handlers.FibHandler)
	healthHandler := http.HandlerFunc(handlers.HealthCheckHandler)

	fibChain := alice.New(
		middlewares.TimeoutMiddleware,
		nosurf.NewPure,
		middlewares.LoggingMiddleware,
	).ThenFunc(fibHandler)

	http.Handle("/fib", fibChain)
	http.Handle("/health-check", healthHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	log.SetPrefix("Fib-API: ")
	log.SetFlags(0)

	port := ":8000"
	fmt.Printf("Starting Webserver on port %v\n", port)

	handleRequests(port)
}
