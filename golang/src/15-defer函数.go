package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename:="file.txt"
	ReadFile(filename)
}

// 读写文件
func ReadFile(filename string)  {
	f1,err:=os.Open(filename)

	// 或者创建一个匿名函数来使用defer
	defer func ()  {
		log.Println("使用匿名函数---准备关闭文件")
		f1.Close()
	}()  // 这里有()是因为调用匿名函数，匿名函数没有名字，属于一次性。类似于lamada表达式

	//defer f1.Close()   // 正常调用

	if err!=nil{
		log.Println("打开文件失败！",err)
		return
	}

	// 如果这里放很多defer，顺序应该是：先文件内容，2222，1111，0000，关闭 ---> 先入后出
	defer fmt.Println("0000")
	defer fmt.Println("1111")
	defer fmt.Println("2222")

	buf:=make([]byte,1024)
	n,err1:=f1.Read(buf)
	if err1!=nil{
		log.Println("读取文件失败！",err1)
		return
	}
	fmt.Println("读取到的内容为:",string(buf[:n]))
}

