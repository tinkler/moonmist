package cjson

import (
	"strings"

	"github.com/tinkler/moonmist/pkg/jsonz"
)

// lower first letter
func ToCamelCase(name string) string {
	if name == "" {
		return ""
	}
	s := jsonz.CommonInitialismsReplacer.Replace(name)
	return strings.ToLower(s[:1]) + s[1:]
}

// SnakeCase to FullCamelCase
func SnakedToFullCamelCase(name string) string {
	if name == "" {
		return ""
	}
	sli := strings.Split(name, "_")
	for i, v := range sli {
		sli[i] = strings.ToUpper(sli[i][:1]) + v[1:]
	}
	return strings.Join(sli, "")
}
