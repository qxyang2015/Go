package struct_demo

import "fmt"

type StructDemo struct {
	Name string //给结构体绑定一个字段，用以说明结构体和结构体指针绑定函数的区别
	Map  map[string]string
}

func NewStructDemo() *StructDemo {
	sd := &StructDemo{
		Name: "newName",
	}
	tmpMap := make(map[string]string)
	sd.Map = tmpMap
	return sd
}

func (sd StructDemo) SetName() {
	sd.Name = "结构体设置Name"
}

func (sd *StructDemo) PointSetName() {
	sd.Name = "结构体指针设置Name"
}

func (sd StructDemo) SetMap() {
	//tM := *(sd.Map)
	sd.Map["1"] = "a"
	fmt.Println("SetMap Print", sd.Map)
}

func (sd *StructDemo) PointSetMap() {
	//tM := *(sd.Map)
	sd.Map["2"] = "B"
	fmt.Println("PointSetMap Print", sd.Map)
}

/*
结构体函数不能实际改变结构体成员变量的值
结构体指针函数可以改变成员变量的值
*/

/*

func main() {
	//调用结构体函数设置Name
	sd := &struct_demo.StructDemo{Name: "初始化Name"}
	sd.SetName()
	fmt.Println("name = ", sd.Name)
	//调用结构体指针函数设置Name
	sd1 := &struct_demo.StructDemo{Name: "初始化Name"}
	sd1.PointSetName()
	fmt.Println("name = ", sd1.Name)

	//
	sd = struct_demo.NewStructDemo()
	sd.SetMap()
	sd.PointSetMap()
	fmt.Println(sd.Map)
}
*/
