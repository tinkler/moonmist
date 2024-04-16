package db

import (
	"context"
	"net/http"

	"github.com/tinkler/moonmist/internal/runtime"
	"gorm.io/gorm"
)

func Use(dbs map[string]*gorm.DB) func(next http.Handler) http.Handler {
	if dbs == nil {
		panic("db is nil")
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValue(r.Context(), runtime.DBContextKey, dbs)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
