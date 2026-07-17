package main

import "fmt"

// 运行到panic之后 后续代码不会执行
// panic 执行之前会 先执行储存的defer
func main() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
	panic("panic")
	defer fmt.Println("D")

}

// 输出结果：
// C
// B
// A
// panic
