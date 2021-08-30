package main

import "fmt"

func main()  {
	p1:=TestPtr()
	fmt.Println("p1的内容是:",*p1)
}

// 如何去查看-p1是存放在栈上还是堆上呢？
// 编译的时候，输入：go build -o test.exe --gcflags "-m -m -l"  11.txt

