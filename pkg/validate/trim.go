package validate

import (
	"reflect"
	"strings"
)

func TrimStructSpaces(s interface{}) {
	// Get the type of the struct
	t := reflect.TypeOf(s)
	// Make sure s is a struct
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return
	}
	// Get the value of the struct
	v := reflect.ValueOf(s).Elem()
	// Iterate over each field of the struct
	for i := 0; i < t.NumField(); i++ {
		// Get the value of the field
		fieldValue := v.Field(i)
		// Check if the field is a string
		if fieldValue.Kind() == reflect.String {
			// Trim spaces from the string value
			trimmedValue := strings.TrimSpace(fieldValue.String())
			// Set the trimmed value back to the field
			fieldValue.SetString(trimmedValue)
		}
	}
}
