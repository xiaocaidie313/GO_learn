package main

import "fmt"

func demoUnbuffered() {
	c := make(chan int)
	go func() {
		fmt.Println("  [无缓冲] 发送 666")
		c <- 666
	}()
	num := <-c
	fmt.Println("  [无缓冲] 收到:", num)
}

func demoBufferedAndSelect() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("  [有缓冲] 读出:", <-c, <-c, <-c)

	ch1, ch2 := make(chan string), make(chan string)
	go func() { ch1 <- "from ch1" }()
	go func() { ch2 <- "from ch2" }()
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("  [select]", msg)
		case msg := <-ch2:
			fmt.Println("  [select]", msg)
		}
	}
}

func main() {
	fmt.Println("=== 无缓冲 channel ===")
	demoUnbuffered()
	fmt.Println("=== 有缓冲 + select ===")
	demoBufferedAndSelect()
}
