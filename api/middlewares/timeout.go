package middlewares

import (
	"net/http"
	"time"
)

// TimeoutMiddleware adds a max time for the server to respond a call.
func TimeoutMiddleware(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 15*time.Second, "Server timed out after 5 seconds\n")
}
