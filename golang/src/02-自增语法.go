package main

import "fmt"

func main()  {
	i:=1
	i++
	fmt.Println("i:",i)
	//fmt.Println("i:",i++)  这是错误的，必须要单独一行

	p:=100
	p--
	fmt.Println("p",p)
}