package middleware

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/sllt/kite-layout/pkg/log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func SignMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	appKey := os.Getenv("API_SIGN_KEY")
	appSecret := os.Getenv("API_SIGN_SECRET")

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}

			for _, header := range requiredHeaders {
				value := r.Header.Get(header)
				if value == "" {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(map[string]any{
						"code":    http.StatusBadRequest,
						"data":    nil,
						"message": "Bad Request",
					})
					return
				}
			}

			data := map[string]string{
				"AppKey":     appKey,
				"Timestamp":  r.Header.Get("Timestamp"),
				"Nonce":      r.Header.Get("Nonce"),
				"AppVersion": r.Header.Get("App-Version"),
			}

			var keys []string
			for k := range data {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

			var str string
			for _, k := range keys {
				str += k + data[k]
			}
			str += appSecret

			if r.Header.Get("Sign") != strings.ToUpper(cryptor.Md5String(str)) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"code":    http.StatusBadRequest,
					"data":    nil,
					"message": "Bad Request",
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
