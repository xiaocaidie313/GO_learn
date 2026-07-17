package main

import (
	"fmt"
	"slices"
)

func main() {
	// 数组 vs 切片
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[1:4] // 共享底层数组
	sli[0] = 99
	fmt.Println("数组:", arr)   // [1 99 3 4 5]
	fmt.Println("切片:", sli)   // [99 3 4]

	// make / append
	s := make([]int, 0, 3)
	s = append(s, 1, 2, 3)
	fmt.Println("append:", s, "len:", len(s), "cap:", cap(s))

	// 克隆（不共享底层数组）
	clone := slices.Clone(s)
	clone[0] = 0
	fmt.Println("原切片:", s, "克隆:", clone)

	// 删除 index=1 的元素
	s = append(s[:1], s[2:]...)
	fmt.Println("delete:", s)
}
