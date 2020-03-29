package main

import (
	"fmt"
	"github.com/qxyang2015/Go/reflect.go"
	"reflect"
)

func main() {
	fmt.Println("start")
	methods := &reflect_go.Methods{
		Age: 18,
	}
	methods.CallMethod("say_hello")
	methods.CallMethod("say_name", reflect.ValueOf("xiao ming"))
	result := methods.CallMethod("get_age")
	fmt.Println(result, len(result), result[0].Interface())
	fmt.Println("done!")
}
