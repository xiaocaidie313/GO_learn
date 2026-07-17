package main

import (
	"fmt"
	"iter"
)

// FibonacciSeq Go 1.23+ 迭代器写法，对应 01.go 的闭包迭代器
func FibonacciSeq(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return // 调用方 break 时，yield 返回 false，提前结束
			}
			a, b = b, a+b
		}
	}
}

func main() {
	fmt.Println("=== Go 1.23+ iter.Seq 写法 ===")
	for v := range FibonacciSeq(10) {
		fmt.Println(v)
	}

	fmt.Println("=== 对比 01.go 闭包写法 ===")
	next := Fibonacci(10)
	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
