package main

import "fmt"

func main() {
	// go 语言也有指针
	// 结构体成员的调用时： C语言中： ptr->name     go语言中: ptr.name

	// go语言在使用指针时，会使用内部的垃圾回收机制（gc:garbage collector，这里所谓的垃圾就是变量内存），开发人员不需要手动释放。
	// C语言需要使用free手动释放

	// c语言不允许返回栈上指针，但是go语言可以返回栈上指针，程序在编译的时候就确定了变量的分配位置
	// 编译的时候，如果发现有必要（有引用），就将变量分配到堆上  （堆：heap, 栈：stack）
	// C语言中，局部变量就是栈上(stack,由操作系统自动分配，存放局部变量等)，如果是new和malloc分配的，就是堆上（heap,由开发人员释放和分配）
	// Go语言中，编译时发现临时使用，放在栈上；如果后续有引用，放在堆上。


	// 1. 使用&定义一个指针
	name:="lily"
	ptr:=&name    // 将name的地址赋予给ptr
	fmt.Println("打印内容：",*ptr)  // lily
	fmt.Println("打印指针：",ptr)   //  0xc000010240

	// 2. 使用new关键字定义指针
	name2Ptr:=new(string)
	* name2Ptr="lily2"
	fmt.Println("打印内容：",* name2Ptr)
	fmt.Println("打印指针：",name2Ptr)

	// 3. 测试 返回栈上指针
	res:=testPtr()
	fmt.Println("打印一下这个栈上指针的值",*res)  //打印一下这个栈上指针的值 ShangHai
	fmt.Println("打印一下这个栈上指针",res)  // 打印一下这个栈上指针的值 ShangHai
	// 可以返回栈上的指针，编译器在编译程序时，会自动判断这段代码，将city变量分配在堆上（heap）

	// 4. 测试 空指针
	// C中空指针：null， C++中空指针：nullptr, golang中空指针：nil
	if res == nil{
		fmt.Println("res是空指针,为nil!")
	}else{
		fmt.Println("res为非空指针！")
	}

}

func testPtr() * string {
	city:="ShangHai"   // 定义了一个局部变量
	ptr:=&city         // 定义了一个指针，并初始化
	return ptr         // 返回一个字符串指针
	// 一般而言，在C语言中，这种是不允许的。局部变量的指针，在函数使用后，局部变量就释放掉了，返回的栈上指针会出错。

}