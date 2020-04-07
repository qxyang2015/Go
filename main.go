package main

import (
	"fmt"
	"github.com/qxyang2015/Go/struct_demo"
)

/*
结构体函数不能实际改变结构体成员变量的值
结构体指针函数可以改变成员变量的值
*/
func main() {
	//调用结构体函数设置Name
	sd := &struct_demo.StructDemo{"初始化Name"}
	sd.SetName()
	fmt.Println("name = ", sd.Name)
	//调用结构体指针函数设置Name
	sd1 := &struct_demo.StructDemo{"初始化Name"}
	sd1.PointSetName()
	fmt.Println("name = ", sd1.Name)
}
