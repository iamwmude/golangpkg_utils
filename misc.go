package utils

import (
	"fmt"
	"reflect"
)

// Put multiple attributes and get first only. This is for testing, do not use in production.
func GetFirstReturn(attributes ...interface{}) interface{} {
	if len(attributes) <= 0 {
		return nil
	}

	return attributes[0]
}

func GetString(v interface{}) string {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Struct:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		if marshalV, err := Marshal(v); err == nil {
			return string(marshalV)
		} else {
			return ""
		}
	default:
		return fmt.Sprintf("%v", v)
	}
}
