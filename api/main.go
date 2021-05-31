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
	// Route setup
	fibHandler := http.HandlerFunc(handlers.FibHandler)
	healthHandler := http.HandlerFunc(handlers.HealthCheckHandler)
	redisHandler := http.HandlerFunc(handlers.RedisHandler)

	// TODO(Frattezi): Can be better encapsulated
	fibChain := alice.New(
		middlewares.TimeoutMiddleware,
		nosurf.NewPure,
		middlewares.LoggingMiddleware,
		middlewares.SetHeadersMiddleware,
	).ThenFunc(fibHandler)

	healthChain := alice.New(
		middlewares.TimeoutMiddleware,
		nosurf.NewPure,
		middlewares.LoggingMiddleware,
		middlewares.SetHeadersMiddleware,
	).ThenFunc(healthHandler)

	// Handlers bind
	http.Handle("/fib", fibChain)
	http.Handle("/health-check", healthChain)
	http.Handle("/redis", redisHandler)

	// Server start
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	log.SetPrefix("Fib-API: ")
	log.SetFlags(0)

	port := ":8000"
	fmt.Printf("Starting Webserver on port %v\n", port)

	handleRequests(port)
}
