package main

import "fmt"

func main() {
	// 1. 定义
	names:=[]string{"北京","上海","广州","深圳"}
	for i:=0;i<len(names);i++{
		fmt.Println(names)
	}

	/*for k,v:=names{
		fmt.Println("key:",k,"value",v)
	}*/

	// 2. 追加数据
	fmt.Println("追加之前的长度:",len(names),"追加之前的容量:",cap(names))  //追加之前的长度: 4 追加之前的容量: 4
	names=append(names,"海南岛")
	// 对于一个切片slice，不仅有'长度'的概念len()，还有一个'容量'的概念cap()
	// 比较追加元素之前和之后，长度 和 容量的变化
	fmt.Println("追加以后的长度:",len(names),"追加以后的容量:",cap(names)) //追加以后的长度: 5 追加以后的容量: 8
	// 出现上述原因是 动态扩容，小于1024，直接翻倍。大于1K，就是1.25倍

	// 3. 基于一个定长数组-->创建一个新的切片slice
	city:=[7]string{"北京","上海","广州","深圳","洛阳","南京","秦皇岛"}
	//city1:=[]string{}    // 创建一个空的slice
	city1:=city[0:2]
	fmt.Println(city1)    // {"北京","上海"},如果修改city1的值，也会修改city的值 ===》引用，切片底层指针指向数组，但是len和cap不变。
	city1[1]="Shanghai"
	fmt.Println("city1:",city1,"city:",city)  // city1: [北京 Shanghai] city: [北京 Shanghai 广州 深圳 洛阳 南京 秦皇岛]

	// 如果想要切片完全独立，可以使用copy()函数来----深拷贝


	// 3. 使用make()函数创建切片 -- 明确切片的长度和容量
	str1:=make([]string,10,20)
	fmt.Println("str1的长度：",len(str1),"str1的容量：",cap(str1))  // 10， 20


	// 深拷贝、浅拷贝
	// 浅拷贝：对于值类型的话是完全拷贝，对于引用类型是拷贝其地址；所以新对象修改的话会影响到原对象值 （对值和地址进行拷贝）常用有：array,int,string,struct,float,bool
	// 深拷贝：对任何对象都会被完整的拷贝一份，相互隔离。（创造一个新对象，不共享内存）常用有：slice,map
	// 如果想让切片完全独立于原始数组，可以使用copy()函数来完成
	nameCopy:=make([]string,cap(names))
	copy(nameCopy,names)    // 如果names是定长数组，使用names[:]来转换成切片
	nameCopy[2]="1111"
	fmt.Println("深拷贝的nameCopy:",nameCopy,"原对象值：",names)  //深拷贝的nameCopy: [北京 上海 1111 深圳 海南岛   ] 原对象值： [北京 上海 广州 深圳 海南岛]


}