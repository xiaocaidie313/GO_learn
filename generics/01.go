package main

import "fmt"

type Int interface {
	~int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64
}

type TinyInt int8

func Do[T Int](n T) T {
	return n + 1
}

func main() {
	// 底层类型在约束集内，可直接使用
	fmt.Println(Do(10))    // 11
	fmt.Println(Do(int64(99))) // 100

	// TinyInt 是自定义类型，不是 int8 本身，需单独约束：
	// Do[TinyInt](1)  // 编译失败
	var t TinyInt = 1
	fmt.Println(TinyInt(Do(int8(t)))) // 需显式转换
}
