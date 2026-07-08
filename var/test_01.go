package main

import (
	"fmt"
)

func main() {

	// 先定义变量
	var b int
	b = 20
	fmt.Println(b)

	var x int = 10
	fmt.Println(x)
	// 系统自己推断类型
	var y = 20
	fmt.Println(y)
	fmt.Printf("y value = %d, y type = %T\n", y, y)
	// 声明变量时 同时赋值  系统自己推断类型
	// 只能用在函数体内  在全局中 只能用前三中
	a := 30
	fmt.Println(a)
	// 多变量声明
	var xx, yy = 100, "5555"
	var (
		kk int    = 200
		ll string = "6666"
	)
	fmt.Println(xx, yy, kk, ll)

	//常量  只读属性
	const length = 100

	// iota  关键字 只有在const 中使用
	// iota 是一个常量计数器  在const关键字出现时被重置为0  const中每新增一行常量声明 iota就会自动加1
	// 之后每一行的表达式都和第一行的 一样 直达出现新的表达式

	const (
		h, i = iota + 1, iota + 2
		c, d // iota = 2 c = iota + 1 = 3, d = iota + 2 = 4
		e, g // iota= 3  e = iota + 1 = 4, g = iota + 2 = 5

		xxx, yyy = iota * 3, iota * 4 // iota = 4  xxx = 12, yyy = 16
		zzz, jjj                      // iota = 5  zzz = 15, jjj = 20
	)

	const (
		A = iota // 0
		B        // 1，自动继承上一行表达式 = iota
		C        // 2
		D = 100  // iota 依然走到3
		E        // 3
	)

}
