package main

import "fmt"

// 规则：人类的武器发起攻击，不同等级子弹效果不同

// 定义一个接口，注意类型是interface
type IAttack interface {
	// 接口函数可以有多个，但是只能有函数原型，不可以有实现
	Attack()
}

// 低等级的类
type HumainLow struct {
	name string
	level int
}

func (this *HumainLow) Attack()  {
	fmt.Println("我是：",this.name,"等级是：",this.level)
}


// 高等级的类
type HumainHigh struct {
	name string
	level int
}

func (this *HumainHigh) Attack()  {
	fmt.Println("我是：",this.name,"等级是：",this.level)
}

// 定义 一个多态的通用的接口：传入不同的对象，调用同样的方法，实现不同的功能。
func  DoAttack(a IAttack)  {
	a.Attack()
}

func main() {
	// 定义一个IAttack接口的对象，任何一个实现该接口的对象，都可以赋值给他。（空接口没用）
	var player IAttack    // var player interface{}属于空接口

	lowLevel:=HumainLow{
		name: "Duke",
		level: 1,
	}
	lowLevel.Attack()

	HighLevel:=HumainLow{
		name: "Duke",
		level: 11,
	}
	lowLevel.Attack()

	// 多态：同一个palyer，传入不同等级的对象；结果不一样
	// 对player赋值为lowLevel，接口需要使用指针类型来赋值
	player=&lowLevel     // 赋值低等级对象
	player.Attack()      // 得到低等级的结果
	player=&HighLevel    // 赋值高等级对象
	player.Attack()      // 得到高等级的结果

	DoAttack(&lowLevel)
	DoAttack(&HighLevel)
	// 能实现主要是因为 低等级和高等级 都实现了Attack的方法。不存在继承关系
}