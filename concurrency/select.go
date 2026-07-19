package main

import (
	"fmt"
	"time"
)

func demoSelectTimeout() {
	fmt.Println("=== select 多路监听 + 超时 ===")
	ch := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "数据到了"
	}()

	select {
	case msg := <-ch:
		fmt.Println("  收到:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("  超时")
	}
}

func demoSelectNilChan() {
	fmt.Println("\n=== nil channel 在 select 中被忽略 ===")
	var nilCh chan int
	select {
	case <-nilCh:
		fmt.Println("  不会走到这")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  只剩 timeout 分支生效")
	}
}

func demoSelectBlockForever() {
	fmt.Println("\n=== select{} 永久阻塞（不占 CPU）===")
	fmt.Println("  类似单片机 WFI 等中断，程序不会退出")
	fmt.Println("  注释掉下面这行可正常结束：")
	// select {}
}

func main() {
	demoSelectTimeout()
	demoSelectNilChan()
	demoSelectBlockForever()
	fmt.Println("\nend")
}
