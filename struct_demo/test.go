package struct_demo

type StructDemo struct {
	Name string //给结构体绑定一个字段，用以说明结构体和结构体指针绑定函数的区别
}

func (sd StructDemo) SetName() {
	sd.Name = "结构体设置Name"
}

func (sd *StructDemo) PointSetName() {
	sd.Name = "结构体指针设置Name"
}
