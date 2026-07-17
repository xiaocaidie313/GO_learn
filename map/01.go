package main

import "fmt"

func main() {
	map2 := map[int]string{
		1: "go",
		2: "c++",
		3: "python",
	}
	fmt.Println("遍历 map:")
	for k, v := range map2 {
		fmt.Printf("  %d → %s\n", k, v)
	}

	delete(map2, 1)
	fmt.Println("delete 后:", map2)
}
