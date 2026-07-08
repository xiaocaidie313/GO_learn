package main

import (
	"fmt"
)

func testarray(mayarray [10]int) {
	fmt.Println("只接受长度一样的数组作为参数")
	// 不同长度的的数组是不同的类型
}

func main() {

	var arrary1 [10]int
	var arraryT [10]int

	var arrary2 [10]int

	for i := 0; i < len(arrary1); i++ {
		arrary1[i] = i + 1
		arraryT[i] = i + 1
	}

	// range 关键字 循环关键字
	// 第一个参数是index 第二个参数是 value
	for index, value := range arraryT {
		arrary2[index] = value
	}

	// 切片  slice
	var s1 []int // 不固定长度的数组就是 切片slice   // 动态数组
	/*
		type slice struct {
			ptr *array   // 指向底层数组
			len int      // 当前长度
			cap int      // 容量
		}
	*/
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)
	s3 := arrary1[0:5] // 切片的使用  取数组的前5个元素
	s3[2] = 999        // slice 是引用类型  修改slice的值 会影响到原来的数组 切出来的是引用

	// 创建slice  make([]T, len, cap)  // cap 可以不写  默认和len一样
	s4 := make([]int, 5)     // [0,0,0,0,0]
	s5 := make([]int, 5, 10) // [0,0,0,0,0]  cap = 10  容量2为10
	fmt.Printf("s5 的长度是 = %d, s5的容量是 %d\n", len(s5), cap(s5))
	// 满了之后自动再扩容一个len长度 比如 3 -> 6
	copy(s4, s5) // copy 复制slice   将s5的全部复制到s4当中

}
