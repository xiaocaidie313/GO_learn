package main

import (
	"fmt"
)

// return 变量 不是单一操作，编译器会拆成固定顺序：
// 第一步：给返回变量赋值
// 第二步：执行所有 defer 延迟函数

// 2-3 之间可以有闭包操作返回值
// 第三步：将返回值传递给上层调用者，函数退出

func foo1() {

	// defer 关键字类似于 析构函数  在函数结束之前的 运行
	// 在函数的生命周期最后
	defer fmt.Println("foo1 finished")
	defer fmt.Println("foo2  finished")
	//存在多个defer的 时收  defer 是 栈的存储  先进后出

}

func main() {
	var a, b int
	a = 1
	b = 2
	defer fmt.Println(sum(a, b)) // defer 会预先 加载函数参数 等到 return 之后执行
	a = 3
	b = 4
	test()
	fmt.Println(test())
}

func sum(a, b int) int {
	return a + b
}
func main2() { // 如果defer 的函数参数是函数 则会在 defer 执行时 计算函数参数的值
	var a, b int
	a = 1
	b = 2
	defer func(num int) {
		fmt.Println(num)
	}(sum(a, b))
	a = 3
	b = 4

	// println() 1+2 = 3
}
func test() (x int) {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	x = 10
	return
	// 输出顺序：defer2 → defer1，再把x返回
}
