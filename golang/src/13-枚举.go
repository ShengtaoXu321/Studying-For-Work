package main

import "fmt"

func main() {
	// 在go语言中没有枚举类型，但是我们可以使用 const+iota(常量累加器) 来进行模拟
	// 模拟一周的枚举
	const (
		MON = iota
		TUE = iota
		WED
		THU
		FRI
		SAT
		SUN
	)
	/*
	iota知识点
	1. iota是常量计数器，从0开始，每换行递增1
	2. 如果不进行赋值，默认与上一行表达式相同
	3. 如果一行出现两个iota，那么两个iota的值是相同的
	4. 每个常量组的iota是独立的，如果遇到const iota会重新清零
	*/

	fmt.Println(SUN)
}