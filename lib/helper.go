package lib

import (
	"log"
	"os"
	"reflect"
	"strings"
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func arrHasKey(arr interface{},key interface{}) {
	if reflect.TypeOf(arr).Name() == "array" {
	}
}
