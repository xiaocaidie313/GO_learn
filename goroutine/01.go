package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask i = ", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine  一旦结束 所有goroutine结束
func main() {

	go newTask()

	i := 0

	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(1 * time.Second)

	}

}
