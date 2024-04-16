package db

import (
	"context"

	"github.com/tinkler/moonmist/internal/runtime"
	"gorm.io/gorm"
)

const DEFAULT_NAME = "default"

func Get(ctx context.Context, v ...interface{}) *gorm.DB {
	if len(v) > 0 {
		if name := getV(v[0]); len(name) > 0 {
			return ctx.Value(runtime.DBContextKey).(map[string]*gorm.DB)[name]
		}
	}
	return ctx.Value(runtime.DBContextKey).(map[string]*gorm.DB)[DEFAULT_NAME]
}

func getV(v interface{}) string {
	s, ok := v.(string)
	if ok {
		return s
	}
	return ""
}
