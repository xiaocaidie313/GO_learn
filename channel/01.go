package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer func() {
			fmt.Println("A defer")
		}()
		fmt.Println("A 开始")

		c <- 666
	}()

	num := <-c
	// 如果新开的 协程和 main 代码不同步 比如 协程先到 c <- 666 则会在协成发生阻塞 直到 main 的 <-c独处
	// 同理如果 <-c先到 则会阻塞  等待 c<-的写入
	fmt.Println("Main 开始, 接受 Num = ", num)
	defer func() {
		fmt.Println("Main defer")
	}()

}
