package main

import "fmt"

func main() {
	// 变量定义：var
	// 常量定义：const

	// 01 - 先定义变量，再赋值
	// 格式 var 变量名 数据类型
	var name string
	var age int
	name = "golang" // 可以使用ctr+alt+l快速格式代码
	fmt.Println("name",name)
	fmt.Println("格式化输出的话")
	fmt.Printf("name is :%s,%d\n",name,age)

	// 02 - 定义时直接赋值
	var gender = "male"
	fmt.Println(gender)

	// 03 - 定义直接赋值，使用自动推导（最常用）
	address:="China"
	fmt.Println(address)

	// 04 - 平行赋值(交叉 )
	i,j:=11,12
	fmt.Println("变化前的i:",i,"j:",j)
	i,i=12,11
	fmt.Println("平行赋值后的i:",i,"j:",j)
}
