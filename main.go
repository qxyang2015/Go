package main

import (
	"fmt"
	"reflect"
)

func Len(v interface{}) int {
	typeVal := reflect.ValueOf(v)

	switch typeVal.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map, reflect.Chan:
		return typeVal.Len()
	default:
		return -1
	}
	return typeVal.Len()
}

func main() {
	fmt.Println("start")
	//[]float32{1.2, 2.5, 36}
	fmt.Println(Len([]string{"1", "2"}))
	fmt.Println("done!")
}
