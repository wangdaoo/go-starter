package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲的 channel
	ch := make(chan int)

	// 启动一个 goroutine，向 channel 发送数据
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("向 channel 发送数据：%d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	// 从 channel 接收数据
	for i := 0; i < 5; i++ {
		data := <-ch
		fmt.Printf("从 channel 接收数据：%d\n", data)
	}
}
