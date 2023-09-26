package acl

import (
	"context"
	"net/http"

	"github.com/tinkler/moonmist/pkg/gs"
	"github.com/tinkler/moonmist/pkg/mst"
)

type contextKey string

const (
	userKey contextKey = "user"
)

func Use(tokenParser func(r *http.Request) (User, error)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, err := tokenParser(r)
			if err != nil {
				gs.HandleError(w, mst.Any(err))
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, u)))
		})
	}
}
