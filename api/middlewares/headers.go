package middlewares

import (
	"net/http"
)

func SetHeadersMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "fib-api")
		w.Header().Set("Content-Type", "application/json")
	}
	return http.HandlerFunc(fn)
}
