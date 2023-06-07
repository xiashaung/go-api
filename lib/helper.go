package lib

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"strings"
)

type helper struct {
}

var (
	Helper = helper{}
)

func (h helper) Error(msg string) error {
	return errors.New(msg)
}

func (h helper) GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (h helper) arrHasKey(arr interface{}, key interface{}) {
	if reflect.TypeOf(arr).Name() == "array" {
	}
}

func (h helper) JsonEncode(value any) string {
	msg, _ := json.Marshal(value)
	return string(msg)
}

func (h helper) JsonDecode(value []byte) any {
	var decode any
	if err := json.Unmarshal(value, &decode); err != nil {
		log.Printf("decode err: %s", err.Error())
	}
	return decode
}
