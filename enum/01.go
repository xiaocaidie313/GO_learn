package main

import "fmt"

type Season int

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	switch s {
	case Spring:
		return "春"
	case Summer:
		return "夏"
	case Autumn:
		return "秋"
	case Winter:
		return "冬"
	default:
		return "未知"
	}
}

type WeekDay string

const (
	Monday    WeekDay = "Monday"
	Tuesday   WeekDay = "Tuesday"
	Wednesday WeekDay = "Wednesday"
)

func main() {
	// iota 枚举
	var s Season = Summer
	fmt.Println("Season:", s, "→", s.String())

	// 字符串枚举
	day := Monday
	fmt.Println("WeekDay:", day)

	// 枚举类型限制：只能赋已定义的常量值
	// s = 99  // 编译可通过但不符合枚举语义，应用自定义类型约束
}
