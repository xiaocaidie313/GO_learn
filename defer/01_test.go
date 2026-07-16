package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	test := func() (x int) {
		defer fmt.Println("defer1")
		defer fmt.Println("defer2")
		x = 10
		return
		// 输出顺序：defer2 → defer1，再把 x 返回
	}
	fmt.Println(test())
}
