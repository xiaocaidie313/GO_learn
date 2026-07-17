package main

import "fmt"

func foo1(a int, b int) (int, int) {
	return a + b, a - b
}

func foo2(a, b int) (res1 int, res2 string) {
	res1 = a + b
	res2 = fmt.Sprintf("a=%d, b=%d", a, b)
	return // 命名返回，裸 return
}

func main() {
	sum, diff := foo1(10, 3)
	fmt.Println("foo1:", sum, diff) // 13 7

	res1, res2 := foo2(5, 2)
	fmt.Println("foo2:", res1, res2) // 7 a=5, b=2
}
