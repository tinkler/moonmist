package db

import (
	"context"

	"github.com/tinkler/moonmist/internal/runtime"
	"gorm.io/gorm"
)

func Get(ctx context.Context, v ...interface{}) *gorm.DB {
	return ctx.Value(runtime.DBContextKey).(*gorm.DB)
}
