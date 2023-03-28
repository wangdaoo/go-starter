package main

import (
	"fmt"
	"sync"
)

// q: 此时 showMsg 并没有完全执行完，为什么程序就已经结束了？
// a: 因为 showMsg 是一个 goroutine，它是并发执行的，不会阻塞主 goroutine 的执行
// q: 如何解决？
// a: 使用 WaitGroup 来等待 showMsg 执行完

// 使用 WaitGroup 来等待 showMsg 执行完
var wg sync.WaitGroup

func showMsg(msg int) {
	// q: wg.Done() 是做什么的？
	// a: wg.Done() 是为了让 wg.Wait() 等待 showMsg 执行完
	defer wg.Done()
	fmt.Println("Message", msg)
}
func main() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		// q: wg.Add(1) 是做什么的？
		// a: wg.Add(1) 是为了让 wg.Wait() 等待 showMsg 执行完
		wg.Add(1)
	}

	// q: wg.Wait() 是做什么的？
	// a: wg.Wait() 是等待 showMsg 执行完，然后再继续执行
	wg.Wait()

	// q: Scanln方法是做什么的？
	// a: Scanln方法是等待用户输入，然后回车，程序才会继续执行
	// fmt.Scanln()
	fmt.Println("Done...")
}
