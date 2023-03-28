package main

import (
	"errors"
	"fmt"
)

func main() {
	// 1. 直接输出错误信息
	err := errors.New("这是一个错误")
	fmt.Println("错误信息：", err)

	// 2. 判断错误是否为 nil
	if err != nil {
		fmt.Println("发生了错误：", err)
	}

	// 3. panic 处理
	// panic 会导致程序崩溃，但是可以通过 recover 恢复
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("程序崩溃了，错误信息：", r)
		}
	}()

	// 4. 返回错误信息
	// 函数返回值中通常会包含一个 error 类型的变量，用于返回错误信息
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("发生了错误：", err)
	} else {
		fmt.Println("结果为：", result)
	}

	/**
	 * panic
	 * panic是一种特殊的错误处理机制，用于表示程序发生了无法处理的错误，必须立即停止程序并打印错误信息
	 */
	panic("这是一个 panic 错误")
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

