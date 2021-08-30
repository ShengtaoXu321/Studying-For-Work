package main

import "fmt"

func main() {
	a,b:=test1(1,2,"abc,")  // 快速看定义:ctr+shift+i
	fmt.Println(a)
	fmt.Println(b)
}

// 函数定义
// 规则：
// 1. 一个函数可以有多个返回值，且返回值同类型可以合并，多个返回值需要使用,分割。
// 2. 定义返回值可以直接加上变量，后续直接使用即可。
func test1(a,b int, c string) (sum int, res string)  {
	sum=a+b
	res=c + "nihao"
	return sum,res  // 这里可以直接写return
}
