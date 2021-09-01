package main

import "fmt"

// 模拟一个父类
type Humain struct {
	name string
	age int
	gender int
}

// 模拟一个继承子类
type Teacher struct {
	Humain   // 继承这里写类型，没有字段名字
	subject string
}

// 结构体嵌套
type Student1 struct {
	hum Humain  // 这个是包含Humain类型的变量，属于类的嵌套
	school string
}

func (this *Humain) Do()  {
	fmt.Println("do something....")
}


func main() {
	t1:=Teacher{}
	t1.name="某老师"
	t1.subject="语文"
	t1.Humain.gender=0   // 这个自动创建了一个默认的同名字段
}