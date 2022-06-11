package server

import (
	"context"
	"net/http"
	"strings"
)

type Token string

// VerifyToken implements a simple middleware handler.
func (s *Server) VerifyToken() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := TokenFromHeader(r)
			newCtx := context.WithValue(ctx, Token("Token"), token)

			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}

// TokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER T".
func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
