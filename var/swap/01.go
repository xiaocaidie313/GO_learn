package swap

import "fmt"

func Compare(a, b int, mode string) int {
	switch mode {
	case "max":
		if a > b {
			return a
		}
		return b
	case "min":
		if a < b {
			return a
		}
		return b
	default:
		return 0
	}
}

func main() {
	a, b := 1, 2
	a, b = b, a // 并行赋值交换
	fmt.Println("交换后:", a, b)

	fmt.Println("max:", Compare(a, b, "max"))
	fmt.Println("min:", Compare(a, b, "min"))
}
