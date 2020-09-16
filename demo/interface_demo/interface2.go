package interface_demo

import (
	"fmt"
	"strconv"
)

type Human2 struct {
	name  string
	age   int
	phone string
}

/*
fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:
type Stringer interface {
	 String() string
}
*/

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human2) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func Demo2() {
	Bob := Human2{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
