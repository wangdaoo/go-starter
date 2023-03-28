package main

import (
	"fmt"
)

func main() {
	// make a slice with capacity of 3
	slice := make([]int, 0, 3)

	// append elements to the slice
	slice = append(slice, 1, 2, 3, 4, 5)

	// Len returns the number of elements in the slice
	fmt.Println("Length of slice:", len(slice)) // Output: 5

	// Cap returns the capacity of the slice
	fmt.Println("Capacity of slice:", cap(slice)) // Output: 6

	// Copy copies elements from source to destination slice, returning the number of elements copied
	newSlice := make([]int, 3)
	numCopied := copy(newSlice, slice)
	fmt.Println("New slice:", newSlice)                  // Output: [1 2 3]
	fmt.Println("Number of elements copied:", numCopied) // Output: 3

	// Cut removes elements from the slice and returns the number of elements removed
	slice = []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("New slice after cut:", slice) // Output: [1 2 4 5]

	// Delete removes an element from the slice by index, preserving the order of the remaining elements
	slice = []int{1, 2, 3, 4, 5}
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("New slice after delete:", slice) // Output: [1 2 4 5]

	// Pop removes the last element from the slice and returns it
	slice = []int{1, 2, 3, 4, 5}
	popped := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	fmt.Println("New slice after pop:", slice) // Output: [1 2 3 4]
	fmt.Println("Popped element:", popped)     // Output: 5

	// Push adds an element to the end of the slice
	slice = []int{1, 2, 3, 4}
	slice = append(slice, 5)
	fmt.Println("New slice after push:", slice) // Output: [1 2 3 4 5]

	// Shift removes the first element from the slice and returns it
	slice = []int{1, 2, 3, 4, 5}
	shifted := slice[0]
	slice = slice[1:]
	fmt.Println("New slice after shift:", slice) // Output: [2 3 4 5]
	fmt.Println("Shifted element:", shifted)     // Output: 1

	// Unshift adds an element to the beginning of the slice
	slice = []int{2, 3, 4, 5}
	slice = append([]int{1}, slice...)
	fmt.Println("New slice after unshift:", slice) // Output: [1 2 3 4 5]

	test()
}

func test() {
	fmt.Println("---test---")

	s1 := make([]int, 3, 5)

	fmt.Println("s1:", s1)

	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	s1 = append(s1, 5)
	fmt.Println("s1:", s1)

	// 创建一个长度为 3，容量为 6 的切片
	s := make([]int, 3, 6)

	s[1] = 1

	// 向切片 s 中追加元素，切片长度和容量都会增加
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)

	// 打印切片 s 的内容
	fmt.Println(s) // [0 1 0 2 3 4 5]

	fmt.Println("-------Cursor-------")

	// slice 内置函数
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println("s2:", s2)

	// 从索引 1 到索引 4（不包含）截取切片 s2
	s2 = s2[1:4]
	fmt.Println("s2[1:4]:", s2) // [2 3 4]

	// 默认下限为 0
	s2 = s2[:2]
	fmt.Println("s2[:2]:", s2) // [2 3]

	// 默认上限为 len(s2)
	s2 = s2[1:]
	fmt.Println("s2[1:]:", s2) // [3]

	// 创建一个切片，长度和容量都是 0
	s2 = make([]int, 0)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=0 cap=0 []

	// 向切片添加一个元素
	s2 = append(s2, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=1 cap=1 [1]

	// 向切片添加多个元素
	s2 = append(s2, 2, 3, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=4 cap=4 [1 2 3 4]

	// 创建一个切片，长度为 3，容量为 5
	s2 = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=3 cap=5 [0 0 0]

	// 向切片添加一个元素
	s2 = append(s2, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=4 cap=5 [0 0 0 1]

	// 向切片添加多个元素，超出容量限制
	s2 = append(s2, []int{2, 3, 4}...)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=7 cap=10 [0 0 0 1 2 3 4]

	// 向切片添加多个元素，超出容量限制
	s2 = append(s2, []int{5, 6, 7}...)
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2) // len=10 cap=10 [0 0 0 1 2 3 4 5 6 7]
}
