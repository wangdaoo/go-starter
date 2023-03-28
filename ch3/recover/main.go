package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到了panic，错误信息为：", err)
		}
	}()

	// 模拟一个panic
	panic("出现了一个错误")
}
