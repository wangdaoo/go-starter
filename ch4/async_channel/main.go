package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个带缓存的 channel，缓存大小为 3
	ch := make(chan int, 3)

	// 开启一个 goroutine，往 channel 中写入数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("写入 %d 到 channel 中\n", i)
		}
	}()

	// 等待 goroutine 写入数据
	time.Sleep(time.Second)

	// 从 channel 中读取数据
	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Printf("从 channel 中读取到 %d\n", value)
		time.Sleep(time.Second)
	}
}
