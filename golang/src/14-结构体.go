package main

import "fmt"

// 1. 定义一个结构体
type Student struct {
	name   string
	age    int
	gender int
	score  float32
}

func main() {
	// 初始化一个结构体
	lily:=Student{
		name:   "lily",
		age:    11,
		gender: 1,
		score:  90,
	}
	fmt.Println("name:",lily.name)

	// 2. 结构体的指针操作
	// Go语言中 指针操作结构体是不用->的，还是使用.
	s1:=&lily
	fmt.Println("使用指针s1:",(*s1).name)
	fmt.Println("使用指针(*s1):",s1.name)

	// 结构体在初始化时，如果对每个字段都赋值，那么字段的名字可以忽略不写。不然，需要明确引用的字段。
}
