package main

import (
	"fmt"
)

func foo1() {

	// defer 关键字类似于 析构函数  在函数结束之前的 运行
	// 在函数的生命周期最后
	defer fmt.Println("foo1 finished")
	defer fmt.Println("foo2  finished")
	//存在多个defer的 时收  defer 是 栈的存储  先进后出

}

func main() {

}
