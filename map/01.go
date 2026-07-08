package main

import (
	"fmt"
)

func main() {
	var map1 map[string]string
	map1 = make(map[string]string, 10)

	for index, value := range map1 {
		fmt.Println(index, value)
	}

	map2 := map[int]string{
		1: "k",
		2: "c++",
		3: "python",
	}
	delete(map2, 1) // 删除某个键
}
