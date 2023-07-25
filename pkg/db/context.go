package db

import (
	"context"
	"net/http"

	"github.com/tinkler/moonmist/internal/runtime"
	"gorm.io/gorm"
)

func Use(db *gorm.DB) func(next http.Handler) http.Handler {
	if db == nil {
		return nil
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValue(r.Context(), runtime.DBContextKey, db)
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
