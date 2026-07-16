package swap

import "fmt"

// max min 函数

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
	a := 1
	b := 2
	a, b = b, a // 直接交换 不需要指针
	fmt.Println(a, b)
	max(a, b, 2.2)
}
