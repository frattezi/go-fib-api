package middlewares

import (
	"net/http"
	"time"
)

func TimeoutMiddleware(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 5*time.Second, "Server timed out after 5 seconds\n")
}
