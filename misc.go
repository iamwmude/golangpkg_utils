package utils

import (
	"errors"
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

func GetMapValue(m interface{}, key interface{}) (interface{}, error) {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		return nil, errors.New("map type error")
	}

	if keys := reflect.ValueOf(m).MapKeys(); len(keys) > 0 && keys[0].Kind() != reflect.TypeOf(key).Kind() {
		return nil, errors.New("key type error")
	}

	for iter := reflect.ValueOf(m).MapRange(); iter.Next(); {
		if iter.Key().Interface() == key {
			return iter.Value().Interface(), nil
		}
	}

	return nil, errors.New("not found")
}
