package lib1

import (
	"fmt"
)

// 首字母大写 代表这是一个对外开放函数
func Add(a int, b int) int {
	return a + b
}

func Lib1Test() {
	fmt.Println("Lib1 test")
}

func init() {
	fmt.Println("lib1 init")
}
