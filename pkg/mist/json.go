package mist

import (
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	json.RegisterExtension(&snakedNamedExtension{})
}

type snakedNamedExtension struct {
	jsoniter.DummyExtension
}

func (ex *snakedNamedExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, field := range structDescriptor.Fields {
		names := make([]string, len(field.ToNames))
		tag, _, _ := strings.Cut(field.Field.Tag().Get("yaml"), ",")
		for i := range field.ToNames {
			if tag != "" {
				names[i] = tag
			}
		}
		field.ToNames = names
		field.FromNames = names
	}
}
