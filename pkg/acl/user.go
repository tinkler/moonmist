package acl

import "context"

type User interface {
	GetUserID() string
	GetRoles() []string
}

func GetUserID(ctx context.Context) string {
	return ctx.Value(userKey).(User).GetUserID()
}

func GetUserRoles(ctx context.Context) []string {
	return ctx.Value(userKey).(User).GetRoles()
}
