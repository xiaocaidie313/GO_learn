package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== goroutine 并发启动 ===")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("  goroutine %d\n", n)
		}(i)
	}
	wg.Wait()

	fmt.Println("\n=== 主 goroutine 退出则程序结束 ===")
	fmt.Println("若不用 WaitGroup/Sleep，子协程可能来不及打印")
	go fmt.Println("  这句可能看不到（主 goroutine 已退出）")
	time.Sleep(100 * time.Millisecond) // 给子协程一点时间
	fmt.Println("end")
}
