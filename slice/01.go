package main

import (
	"fmt"
	"slices"
)

func testarray(mayarray [10]int) {
	fmt.Println("只接受长度一样的数组作为参数")
	// 不同长度的的数组是不同的类型
}

func main() {

	// 数组
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

func ArrayToSlice(array [10]int) []int {
	// [a:b]  左闭右开
	// 全范围就是 转换成slice了

	// 由array 转化的slice 和 array共享地址  slice的转化是 引用类型 修改slice的值 会影响到array的值
	return array[:]
}

func NewSlice(array [10]int) {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := slices.Clone(arr[:]) // 开新空间 克隆
	slice[0] = 0
	fmt.Printf("array: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
}
func Insert() {

	// append 方法本质就只能从尾部插入
	// func append(slice []Type, elems ...Type) []Type
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums = append([]int{-1, 0}, nums...)
	fmt.Println(nums) // [-1 0 1 2 3 4 5 6 7 8 9 10]
	var i int = 3     // 插入位置
	nums = append(nums[:i+1], append([]int{999, 999}, nums[i+1:]...)...)
	fmt.Println(nums) // i=3，[1 2 3 4 999 999 5 6 7 8 9 10]
	nums = append(nums, 99, 100)
	fmt.Println(nums) // [1 2 3 4 5 6 7 8 9 10 99 100]
}

func Delete() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int = 3 // 删除位置
	nums = append(nums[:i], nums[i+1:]...)
	fmt.Println(nums) // [1 2 3 5 6 7 8 9 10]

	var n int = 3 // 删除前n个元素
	nums = nums[n:]
	fmt.Println(nums) //n=3 [4 5 6 7 8 9 10]

	// 删除所有元素
	nums = nums[:0]
	fmt.Println(nums) // []
}

func COPY() {
	dest := make([]int, 0)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src)) // copy的时候要保证有相同的长度
	fmt.Println(src, dest)       // [1 2 3 4 5 6 7 8 9] []
}

func CapAndAppend() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 底层数组：[1,2,3,4,5,6,7,8,9]
	// s1 指针=下标0，len=9，cap=9

	s2 := s1[3:4]
	// 切片语法 s[a:b]：取下标 a ~ b-1
	// s2 指向底层数组下标3，只截取1个元素：[4]
	// s2 的 len=1，cap = 原切片cap - 起始偏移 = 9-3 = 6
	// s2 能访问到底层数组下标 3、4、5、6、7、8 一共6个位置

	// append 方法本质就只能从切片尾部“写入“ 在如果是容量足够的情况下  会直接修改掉切片一下一个元素的地址 也就是原数组的地址  因为切片是能访问到这个为位置的
	// slice切片 底层数组不变  但切片可变长  slice 本身只是个结构体
	fmt.Println(s1, s2)
}
