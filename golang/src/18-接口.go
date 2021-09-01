package main

import "fmt"

func main() {
	// 1. 接口的定义
	// 定义三个接口，分别赋值不同类型的数据
	var i, j, k interface{}
	name := []string{"lily", "duke"}
	age := 10
	school := "中学"
	i = name   // 切片
	j = age    // int
	k = school // string
	fmt.Println("存放切片类型：", i, "存放int类型：", j, "存放字符串类型：", k)

	// 2. 判断接口的类型 ---> 类比于 哈希表的value,ok判断方式
	kvalue, ok := k.(int)
	if !ok {
		fmt.Println("k不是int类型")
	} else {
		fmt.Println("k是int类型，值为：", kvalue)
	}

	// 3. 将interface当成一个函数的参数（类似于print），使用switch来判断用户输入的不同类型。根据不同类型，做相应的逻辑处理
	// 创建一个具有三个接口类型的切片
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello world"
	array[2] = true

	for _, value := range array {
		// 获取当前接口的真正的数据类型
		switch v := value.(type) {
		case int:
			fmt.Println("是int类型的,内容是：",v)
		case string:
			fmt.Println("是字符串类型，内容是：",v)
		case bool:
			fmt.Println("是bool类型，内容是：",v)
		default:
			fmt.Println("不是上述存在类型")
		}
	}

}
