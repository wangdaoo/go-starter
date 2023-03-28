package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前 Goroutine 的数量
	fmt.Println("当前 Goroutine 的数量：", runtime.NumGoroutine())

	// 获取当前操作系统的 CPU 核心数
	fmt.Println("当前操作系统的 CPU 核心数：", runtime.NumCPU())

	// 设置当前 Goroutine 的最大可同时执行数量
	runtime.GOMAXPROCS(2)

	// 获取当前 Goroutine 的最大可同时执行数量
	fmt.Println("当前 Goroutine 的最大可同时执行数量：", runtime.GOMAXPROCS(0))

	// 获取当前程序的内存使用情况
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("当前程序的内存使用情况：\nAlloc：%d bytes\nTotalAlloc：%d bytes\nSys：%d bytes\nNumGC：%d\n", mem.Alloc, mem.TotalAlloc, mem.Sys, mem.NumGC)

	// 获取 goroot 目录
	fmt.Println("goroot 目录：", runtime.GOROOT())

	// 获取操作系统
	fmt.Println("操作系统：", runtime.GOOS)
}
