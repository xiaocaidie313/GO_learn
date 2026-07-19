package main

import (
	"fmt"
	"time"
)

func demoSelectBlock() {
	fmt.Println("=== select{} 永久阻塞 ===")
	fmt.Println("start")
	// select {} // 永久阻塞，类似单片机 WFI 等中断，不占 CPU
	// 取消注释上一行后，下面 end 不会打印
	fmt.Println("end（select{} 已注释，程序可正常结束）")
}

func demoSelectNilAndTimeout() {
	fmt.Println("\n=== select 多路监听 ===")
	var nilCh chan int // 管道为 nil  在select中 被忽略
	select {
	case <-nilCh:
		fmt.Println("read") // 不会走到这
	case nilCh <- 1:
		fmt.Println("write") // 不会走到这
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func main() {
	demoSelectBlock()
	demoSelectNilAndTimeout()
}
