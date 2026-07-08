package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	
	x :=10
	y :=10
	// 匿名函数 
	go func (a, b int) int {
		return a+b
	}(x, y)

	go func(){


		defer fmt.Println("A defer")
		fmt.Println("A")

		go func(){
			defer fmt.Println("B defer")
			// 退出 go 协程
			runtime.Goexit()
			fmt.Println("B")
		}()// 不加 () 就是只是单纯的 声明 这个匿名函数 加上后相当于直接 执行
	}()

	for {
		time.Sleep(1 * time.Second))
	}
}