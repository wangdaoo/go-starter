package main

import (
	"fmt"
	"math"
)

func main() {
	// append()函数用于向切片中追加元素
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	// copy()函数用于将一个切片复制到另一个切片中
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice2)

	// len()函数用于获取切片或数组的长度
	array := [3]int{1, 2, 3}
	fmt.Println(len(array))

	// cap()函数用于获取切片的容量
	slice3 := make([]int, 3, 5)
	fmt.Println(cap(slice3))

	// make()函数用于创建切片、映射和通道
	slice4 := make([]int, 3)
	fmt.Println(slice4)

	// new()函数用于创建指针类型的变量
	ptr := new(int)
	fmt.Println(*ptr)

	// math包中的常用函数
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Ceil(1.5))
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(4))
}
