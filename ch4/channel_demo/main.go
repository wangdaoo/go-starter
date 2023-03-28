package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // 创建一个字符串类型的通道

	// go func() 是一个匿名函数，它会在一个新的 goroutine 中执行
	// 该匿名函数会向通道中写入一个字符串
	go func() {
		ch <- "Hello from channel"
	}()

	// 主 goroutine 会等待 1 秒钟，然后从通道中读取一个字符串
	// 此时的 time.Sleep() 是为了让匿名函数有足够的时间向通道中写入一个字符串
	time.Sleep(time.Second * 1)

	// 从通道中读取一个字符串
	msg := <-ch
	fmt.Println(msg)
}
