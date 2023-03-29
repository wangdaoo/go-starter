package main

import (
	"bufio"
	"fmt"
	"os"
)

// q: Scanner 是什么？
// a: Scanner 是一个读取器，用于读取文件中的数据, 一行一行的读取

func main() {
	// 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 创建一个新的Scanner对象
	scanner := bufio.NewScanner(file)

	// 循环读取文件中的每一行
	for scanner.Scan() {
		// 输出每一行的内容
		fmt.Println(scanner.Text())
	}

	// 检查是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("文件读取错误", err)
	}
}
