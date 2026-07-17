package main

import (
	"fmt"
	"os"
)

func main() {
	// 演示 ReadDir（不依赖 test.txt 存在）
	fmt.Println("当前目录文件:")
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, e := range entries {
		fmt.Println(" ", e.Name())
	}

	// 文件打开示例（test.txt 不存在时走错误分支）
	file, err := os.Open("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("test.txt 不存在（正常演示）")
		return
	}
	if err != nil {
		fmt.Println("打开异常:", err)
		return
	}
	defer file.Close()
	fmt.Println("文件打开成功")
}
