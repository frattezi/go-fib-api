package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frattezi/go-fib-api/handlers"
	"github.com/frattezi/go-fib-api/helpers"
)

// handlerResquests will setup routes and start an http server for a given port.
func handleRequests(port string) {
	fibHandler := http.HandlerFunc(handlers.FibHandler)
	fibRedisHandler := http.HandlerFunc(handlers.FibRedisHandler)
	healthHandler := http.HandlerFunc(handlers.HealthCheckHandler)

	fibChain := helpers.ChainMiddlewares(fibHandler)
	fibRedisChain := helpers.ChainMiddlewares(fibRedisHandler)
	healthChain := helpers.ChainMiddlewares(healthHandler)

	http.Handle("/fib", fibChain)
	http.Handle("/fibr", fibRedisChain)
	http.Handle("/health-check", healthChain)

	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	log.SetPrefix("Fib-API: ")
	log.SetFlags(0)

	port := ":8000"
	fmt.Printf("Starting Webserver on port %v\n", port)

	handleRequests(port)
}
