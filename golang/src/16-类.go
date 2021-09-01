package main

import "fmt"

// 使用struct模拟类
// C++中使用pubilc、private来区别不同权限，Go通过大小写来区分。首字母大写为pubilc，小写为private
// C++的方法是在内部，Go在类的外面绑定方法

type Person struct {
	name   string
	age    int
	gender int
	score  float32
}

// 在类的外面绑定方法
func (p Person) Eat() { // 如果是p，跟当前类就是拷贝的关系
	fmt.Println("吃的方法！")
}

func (this *Person) Eat1() { // 也可以是this指针
	// 类的方法，可以使用自己的成员
	this.name="duke"
	//fmt.Println(this.name)
}

func (this Person) Eat2() { // 不是指针，是一个值得拷贝
	this.name = "duke"
	//fmt.Println(this.name)
}
func main() {
	lily := Person{
		name:   "lily",
		age:    18,
		gender: 0,
		score:  100,
	}
	lily1:=lily  // 防止指针修改干扰原值

	fmt.Println("使用this指针, 修改name值...")
	fmt.Println("修改前的值：",lily)  // {lily 18 0 100}

	lily.Eat1()   // duke
	fmt.Println("修改后的值：",lily)  // {duke 18 0 100}

	fmt.Println("----------------------")
	fmt.Println("不使用this指针，修改name值...")
	fmt.Println("修改前的值：",lily1) //  {lily 18 0 100}
	lily.Eat2()  // duke
	fmt.Println("修改后的值：",lily1) //  {lily 18 0 100}
}
