package server

import (
	"net/http"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
)

func RateLimiter(lmt *limiter.Limiter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// If master key set, no rate limit.
			mk := r.Header.Get("Master-Key")
			if mk != "" {
				next.ServeHTTP(w, r)
				return
			}

			// Get auth token and limit by token.
			token := TokenFromHeader(r)
			err := tollbooth.LimitByKeys(lmt, []string{token})
			if err != nil {
				data := struct {
					Error string
				}{
					"Request limit reached, naughty naughty!",
				}
				writeJSON(w, data, err.StatusCode)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
