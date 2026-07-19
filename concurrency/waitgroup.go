package main

import (
	"fmt"
	"sync"
)

// func (wg *WaitGroup) Add(delta int)
// func (wg *WaitGroup) Done()
// func (wg *WaitGroup) Wait()

// waitGroup就三个 方法 Add Done Wait  分别是计数器
// Add 添加计数器 每次调用Add方法，计数器值增加delta
// Done 完成计数器 每次调用Done方法，计数器值减少1
// Wait 等待计数器 每次调用Wait方法，如果计数器值大于0，则阻塞等待，直到计数器值为0
func main() {
	var wait sync.WaitGroup
	// 指定子协程的数量
	wait.Add(1)
	go func() {
		fmt.Println(1)
		// 执行完毕
		wait.Done()
	}()
	// 等待子协程
	wait.Wait()
	fmt.Println(2)
}
