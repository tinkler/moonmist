package cjson

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tinkler/moonmist/pkg/jsonz"
)

type camelCasedNamedExtension struct {
	jsoniter.DummyExtension
}

func (ex *camelCasedNamedExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, field := range structDescriptor.Fields {
		if jt, ok := field.Field.Tag().Lookup("json"); ok {
			n, _ := jsonz.ParseTag(jt)
			if n != "" {
				continue
			}
		}
		newNames := make([]string, len(field.ToNames))
		for i := range field.ToNames {
			newNames[i] = ToCamelCase(field.ToNames[i])
		}
		field.ToNames = newNames
		field.FromNames = newNames
	}
}
