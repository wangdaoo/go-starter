package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建两个 channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 向 ch1 中发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Second * 1)
		}
	}()

	// 向 ch2 中发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i
			time.Sleep(time.Second * 2)
		}
	}()

	// 使用 select 语句进行多路复用
	for i := 0; i < 10; i++ {
		select {
		case num := <-ch1:
			fmt.Println("从 ch1 中读取到数据：", num)
		case num := <-ch2:
			fmt.Println("从 ch2 中读取到数据：", num)
		case <-time.After(time.Second * 3): // 3 秒后从 time.After() 中读取数据
			fmt.Println("从 ch1 和 ch2 中读取数据超时")
		}
	}
}
