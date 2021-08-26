package main

import "fmt"

/* 题目描述：两个数组的交集
给定两个数组，编写一个函数来计算他们的交集

示例1： num1=[1,2,2,1], num2=[2,2]
输出：[2,2]

示例2： num1=[4,9,5], num2=[9,4,9,8,4]
输出:[4,9]

说明：
	1. 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致
	2. 不考虑输出的顺序
思路：
	设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组。
分析：
	看成传统的映射题（map映射），原因是 需要找出两个数组的交集元素，并且还满足与两数组中出现的次数一致。======>我们需要知道每个值出现的次数
	形成一个映射关系<元素，出现的次数>
*/




func main()  {
	//var n1,n2 int
	//fmt.Print("请输入数组num1的长度n1：")
	//fmt.Scanf("%d",&n1)
	//fmt.Println("请输入长度n2：")
	//fmt.Scanln(&n2)
	//var num1 = make([]int,n1)
	//var num2=make([]int,n2)
	//for i:=0;i<n1;i++{
	//	fmt.Println("请输入num1的值")
	//	fmt.Scanln(&num1[i])
	//}
	//for i:=0;i<n2;i++{
	//	fmt.Println("请输入num2的值")
	//	fmt.Scanln(&num2[i])
	//}
	//fmt.Println(num1)
	//fmt.Println(num2)
	var num1=[]int{4,9,5}
	var num2=[]int{9,4,9,8,4}
	s1:=intersect(num1,num2)
	fmt.Println(s1)

}

func intersect(nums1 []int, nums2 []int) []int{
	//创建一个key为int,值为int类型的空map
	m0 := map[int]int{}
	//m0:=make(map[int]int，0)
	for _,v:=range nums1{
		// 遍历nums1，初始化map
		//fmt.Println("输出的",m0[v])
		m0[v] += 1
		//fmt.Println("输出的",m0[v])
	}
	k:=0
	for _,v:=range nums2{
		//如果元素相同，将其存入nums2中，并将出现的次数减1
		if m0[v]>0{
			m0[v] -= 1
			nums2[k]=v
			k++
		}
	}
	return nums2[0:k]
}
