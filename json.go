package utils

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

func Marshal(v interface{}) ([]byte, error) {
	val, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
	if err != nil {
		log.Printf("marshal failed. reason: %s\n", err.Error())
	}

	return val, err
}

func Unmarshal(data []byte, v interface{}) error {
	err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(data, v)
	if err != nil {
		log.Printf("unmarshal failed. reason: %s\n", err.Error())
	}

	return err
}
