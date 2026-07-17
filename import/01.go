package main

import (
	"fmt"
	. "demo/example"
	t "time"
)

func main() {
	fmt.Println("time:", t.Now().Format("15:04:05"))
	SayHi() // 点导入，无需 example. 前缀
}
