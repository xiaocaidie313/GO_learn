package main

//
import (
	"fmt"
)

func sum(nums ...int) int {
	total := 0
	for _, t := range nums {
		total += t
	}
	return total
}
func main() {
	numberArray := []int{1, 3, 5, 7, 9}
	var total int = sum(numberArray...)
	//  number array被展开成单独的参数传递给sum函数
	fmt.Println(total)

	for i, c := range "hello" {
		fmt.Printf("index: %d, character: %c\n", i, c)
	}
}
