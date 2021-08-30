package main

import "fmt"

func main() {
	// 1. 定义定长数组
	// c语言中定义数组：int arr[10]={1,2,3,4}
	// go语言中定义数组： var arr [10]int = [10]int{1,2,3,4} 或者 var arr = [10]int{1,2,3,4}  最常用：arr:=[10]int{1,2,3,4}
	num:=[10]int{1,2,3,4,5,6}

	// 2. 数组遍历
	// 方法一
	for i:=0;i<len(num);i++{
		fmt.Println("i:",i,"v:",num[i])
	}

	// 方法二：for range
	for key,value:=range num{
		// 常见错误点：修改value值，num数组对应的值不会被修改。 ===》value值只是一个副本，全程只是一个临时变量，不断被重新赋值
		value+=1
		fmt.Println("i:",key,"原数组的值:",num[key])   // 数组值不变
		fmt.Println("i:",key,"修改过后的值:",value)    // value的每个值加一
		// 使用:= 自动推导，必须有新变量
	}

}
