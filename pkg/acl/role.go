package acl

import "context"

// builtin roles
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type CheckType int

const (
	_ CheckType = iota - 1
	// check type full, all roles match
	CT_ONE
	// check type one, one of roles match
	CT_FULL
)

func HasRole(ctx context.Context, t CheckType, roles ...string) bool {
	userInter := ctx.Value(userKey)
	user, ok := userInter.(User)
	if !ok {
		return false
	}
	switch t {
	case CT_FULL:
		roleSet := make(map[string]struct{}, len(roles))
		for _, r := range roles {
			roleSet[r] = struct{}{}
		}
		for _, target := range user.GetRoles() {
			if _, ok := roleSet[target]; !ok {
				return false
			}
		}
		return true
	case CT_ONE:
		for _, target := range roles {
			for _, r := range user.GetRoles() {
				if target == r {
					return true
				}
			}
		}
		return false
	default:
		return false
	}
}
