package middleware

import (
	"net/http"
)

func CORSMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == "OPTIONS" {
				requestMethod := r.Header.Get("Access-Control-Request-Method")
				requestHeaders := r.Header.Get("Access-Control-Request-Headers")
				w.Header().Set("Access-Control-Allow-Methods", requestMethod)
				w.Header().Set("Access-Control-Allow-Headers", requestHeaders)
				w.Header().Set("Access-Control-Max-Age", "7200")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
