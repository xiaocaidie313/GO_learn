package main

import (
	"fmt"
	"sync"
	"time"
)

// 主 goroutine  一旦结束 所有goroutine结束
func main() {
	fmt.Println("start")

	// 开了10个协程  每个协程运行时间不一样
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)
	}
	wg.Wait()

	// 若不用 WaitGroup，主 goroutine 先退出，子协程可能来不及执行
	go fmt.Println("这句可能看不到（主 goroutine 已退出）")
	// 暂停1ms
	time.Sleep(time.Millisecond)
	fmt.Println("end")
}
