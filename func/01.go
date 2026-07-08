package main

import (
	"fmt"
)

// 函数反回值  可以有多个返回值
// 匿名返回
func foo1(a int, b int) (int, int) {
	return a + b, a - b
}

// 命名返回
func foo2(a, b int) (res1 int, res2 string) {
	res1 = a + b
	res2 = fmt.Sprintf("a=%d, b=%d", a, b)

	// 隐式 返回
	return
}
func foo3(a int, b int) int {

	return a + b

}

func main() {
	fmt.Println("Hello, World!")
}
