package main

import (
	"fmt"
	"runtime"
)

// https://pkg.go.dev/runtime

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Message -> ", msg)
	}
	// runtime.Goexit() 是退出当前 goroutine
}

func main() {
	go showMsg("Hello")

	for i := 0; i < 10; i++ {
		// q: runtime.Gosched() 是做什么的？
		// a: runtime.Gosched() 是让出 CPU 时间片，让其他 goroutine 执行
		runtime.Gosched()
		fmt.Println("Main -> ", i)
	}
	fmt.Println("Main ->", "Done...")
}
