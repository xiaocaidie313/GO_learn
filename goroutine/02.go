package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	x, y := 10, 10

	go func(a, b int) {
		fmt.Println("匿名协程:", a+b)
	}(x, y)

	go func() {
		defer fmt.Println("A defer")
		fmt.Println("A")

		go func() {
			defer fmt.Println("B defer")
			runtime.Goexit() // 只退出当前 goroutine
			fmt.Println("B") // 不会执行
		}()
	}()

	time.Sleep(2 * time.Second)
}
