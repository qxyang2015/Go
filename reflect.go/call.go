package reflect_go

import (
	"fmt"
	"github.com/qxyang2015/Go/utils"
	"reflect"
)

/*
	reflect实现的一种，通过函数名统一调用函数的方式
*/

type Methods struct {
	Name string
	Age  int
}

func (methods Methods) CallMethod(name string, params ...reflect.Value) []reflect.Value {
	methodsRef := reflect.ValueOf(methods)

	methodName := utils.Camel(name)

	providerMethod := methodsRef.MethodByName(methodName)
	result := providerMethod.Call(params)
	return result
}

func (methods Methods) SayHello() {
	fmt.Println("hello world!")
}

func (methods Methods) SayName(name string) {
	fmt.Println("my name is:", name)
}

func (methods Methods) GetAge() (int, int) {
	return methods.Age, methods.Age + 1
}

/*
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
*/
