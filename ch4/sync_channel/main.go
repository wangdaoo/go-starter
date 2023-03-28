package main

import (
	"fmt"
	"time"
)

// q: 为什么这里的 channel 是无缓存的？
// a: 因为我们需要保证主 goroutine 在接收到数据之前，匿名函数已经将数据写入到 channel 中了。
//
//	如果 channel 是有缓存的，那么匿名函数可能会在主 goroutine 接收数据之前就将数据写入到 channel 中，
//	这样主 goroutine 就会一直阻塞在接收数据的那一行代码上，从而导致程序无法正常退出。
//
// 无缓存的 channel 会强制匿名函数在向 channel 发送数据之前，必须等待主 goroutine 从 channel 中接收数据。
// 无缓存 channel 是一个**同步**的 channel，它会强制数据的发送和接收在同一时刻发生。
// https://gopl-zh.github.io/ch8/ch8-04.html#841-%E4%B8%8D%E5%B8%A6%E7%BC%93%E5%AD%98%E7%9A%84channels
func main() {
	// 创建一个无缓存的 channel
	ch := make(chan int)

	// 开启一个 goroutine，向 channel 发送数据
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
