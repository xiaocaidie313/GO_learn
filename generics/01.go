package main

type Int interface {
	~int8 | int16 | int | int32 | int64 | uint8 | uint16 | uint | uint32 | uint64
}

// ~int8  表示底层类型是 int8

type TinyInt int8

func Do[T Int](n T) T {
	return n
}

func main() {
	Do[TinyInt](1) // 无法通过编译，即便其底层类型属于Int类型集的范围内
}
