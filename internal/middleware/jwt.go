package middleware

import (
	"context"
	"encoding/json"
	"github.com/sllt/kite-layout/pkg/jwt"
	"github.com/sllt/kite-layout/pkg/log"
	"net/http"
)

type contextKey string

const ClaimsKey contextKey = "claims"

func StrictAuth(j *jwt.JWT, logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				logger.Warnf("No token url=%s", r.URL)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"code":    http.StatusUnauthorized,
					"data":    nil,
					"message": "Unauthorized",
				})
				return
			}

			claims, err := j.ParseToken(tokenString)
			if err != nil {
				logger.Errorf("token error url=%s err=%v", r.URL, err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"code":    http.StatusUnauthorized,
					"data":    nil,
					"message": "Unauthorized",
				})
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func NoStrictAuth(j *jwt.JWT, logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				cookie, err := r.Cookie("accessToken")
				if err == nil {
					tokenString = cookie.Value
				}
			}
			if tokenString == "" {
				tokenString = r.URL.Query().Get("accessToken")
			}
			if tokenString == "" {
				next.ServeHTTP(w, r)
				return
			}

			claims, err := j.ParseToken(tokenString)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
