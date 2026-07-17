package main

import "fmt"

type TwoDMap map[string]map[string]int

func (t TwoDMap) Count() int {
	total := 0
	for _, inner := range t {
		total += len(inner)
	}
	return total
}

type TwoDMap2 = map[string]map[string]int

func main() {
	// 类型定义：可绑方法
	m1 := TwoDMap{"北京": {"朝阳": 1, "海淀": 2}}
	fmt.Println("TwoDMap Count:", m1.Count())

	// 类型别名：等价于原生 map
	var m2 TwoDMap2 = make(map[string]map[string]int)
	m2["上海"] = map[string]int{"浦东": 3}
	fmt.Println("TwoDMap2:", m2)

	// 类型断言
	var a interface{} = 42
	if v, ok := a.(int); ok {
		fmt.Println("断言成功, int =", v)
	}
}
