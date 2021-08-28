package main

import "fmt"

func main() {

	// 1. 字符串的定义
	name := "lily"
	fmt.Println("name:", name)

	// 2. 原生输出
	// 在go中原生输出字符串，使用``来包裹，且里面需要换行
	// 在Python中使用''' aa '''，三个引号来处理
	// 在c语言中使用 单引号' + \ 来解决
	usage := `./a.out <option>
	-h   help
	-a   all`
	fmt.Println("字符串原生输出为：",usage)

	// 3. 字符串长度
	// C++ ： 使用name.length来查看字符串长度
	// Go: string没有.length方法，可以使用自有函数len()进行处理
	// len: 很常用的
	l1:= len(name)
	fmt.Println("字符串name的长度为：",l1)


	// 4. 逐字输出字符串的单个字符
	for i:=0;i<len(name);i++{
		fmt.Printf("i:%d,v:%c\n",i,name[i])
	}

	// 5. 字符串的拼接
	i,j:="hello","world"
	fmt.Println("字符串i拼接j后：",i+j)  // 字符串i拼接j后： helloworld
}
