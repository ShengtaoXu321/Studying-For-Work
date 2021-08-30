package main

import (
	"fmt"
	"os"
)

// switch应用：在命令行输入参数。在switch中进行处理
// 输入参数： C中：argc, **argv     Go中：os.Args  ==》 直接可以获取命令输入，是一个字符串切片
func main() {
	cmds:=os.Args
	// 通过对cmds这个切片的遍历，发现了cmds[0]存放是文件的名字，cmds[1]是输入的一个命令
	// 逻辑判断，如果输入的命令大小不对，就退出
	if len(cmds)<2{
		fmt.Println("请输入正确的参数！")
		return
	}

	// Go中的switch与C语言的区别：默认加上了break，不需要手动处理
	// 如果想向下穿透，需要加上关键字:fallthrough
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		fallthrough  // 向下穿透
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("dealut")
	}

}
