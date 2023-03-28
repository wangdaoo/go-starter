package main

import (
	"fmt"
)

func main() {
	// 创建一个整型类型的channel
	intChan := make(chan int)
	// 创建一个字符串类型的channel
	strChan := make(chan string)

	// 启动一个goroutine，向intChan中写入整型数据
	go func() {
		for i := 0; i < 5; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	// 启动一个goroutine，向strChan中写入字符串数据
	go func() {
		for i := 0; i < 5; i++ {
			strChan <- fmt.Sprintf("string %d", i)
		}
		close(strChan)
	}()

	// 启动一个goroutine，从intChan中读取整型数据，并将其转换为字符串后写入到strChan中
	go func() {
		for i := range intChan {
			strChan <- fmt.Sprintf("int %d", i)
		}
		close(strChan)
	}()

	// 从strChan中读取数据并打印
	for str := range strChan {
		fmt.Println(str)
	}
}
