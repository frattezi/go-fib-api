package helpers

import (
	"net/http"

	"github.com/frattezi/go-fib-api/middlewares"
	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
)

// ChainMiddlewares creates a new chain of custom middlewares
func ChainMiddlewares(handler http.HandlerFunc) http.Handler {
	// TODO(Frattezi): Enhance factory for custom setups.
	customHandler := alice.New(
		middlewares.TimeoutMiddleware,
		nosurf.NewPure,
		middlewares.LoggingMiddleware,
	).ThenFunc(handler)

	return customHandler
}
