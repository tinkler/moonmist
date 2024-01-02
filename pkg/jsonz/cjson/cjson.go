package cjson

import (
	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

func init() {
	Json.RegisterExtension(&camelCasedNamedExtension{})
}

func Unmarshal(data []byte, v any) error {
	return Json.Unmarshal(data, v)
}

func Marshal(v any) ([]byte, error) {
	return Json.Marshal(v)
}
