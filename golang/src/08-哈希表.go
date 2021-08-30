package main

import "fmt"

func main() {
	// 1. 定义map
	// 定义内容是：key:学生id   value:学生名字
	var idNams map[int]string // 注意：使用这种方式定义完以后是不能直接赋值的，因为是空的
	// 使用map之前，一定要分配空间
	idNams = make(map[int]string, 10) // 使用map分配空间,可以不指定长度
	idNams[0] = "duke"
	idNams[1] = "edsion"
	for key, value := range idNams {
		fmt.Println("key:", key, "value:", value)
	}

	// 2. 使用自动推导定义并且直接分配空间
	//idNams:=make(map[int]string,10)

	// 3. 判断一个key是否存在map中  ==》 value,ok机制
	name9 := idNams[8]
	fmt.Println("name9是否存在：", name9) // 空空如也，因为是字符串类型，所以返回空。如果是int类型，返回0
	// 总结：在map中不存在访问越界问题。所有的key都是唯一且有效的，对同一个key赋值就是覆盖；访问一个不存在的key不会崩溃，返回这个类型的零值
	// 零值：bool==> false  数字==>0  字符串==>空
	// 所以无法通过value值来判断一个key值是否存在，因为可能存在和不存在就是空。 ===》 value,ok机制
	value, ok := idNams[1]
	if ok {  // 如果id=1是存在的，那么value的值就是key=1对应的值；ok返回true。否则就是返回空值和False
		fmt.Println("id=2这个值存在，值为",value)
	}else{
		fmt.Println("不存在！")
	}


	// 4. 删除map中的元素
	// 使用delete()来删除指定的key
	delete(idNams,1)
	fmt.Println("删除后的map：",idNams)
	delete(idNams,9)  // 删除一个没有的，也不会报错

	// 5. 并发处理时，需要对map进行上锁
}
