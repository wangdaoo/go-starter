package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	count int
	wg    sync.WaitGroup
	rwMtx sync.RWMutex
)

func main() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}
	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}

func read(n int) {
	defer wg.Done()
	rwMtx.RLock() // 读锁定
	fmt.Printf("读goroutine %d 正在读取...\n", n)
	v := count
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	rwMtx.RUnlock() // 读解锁
}

func write(n int) {
	defer wg.Done()
	rwMtx.Lock() // 写锁定
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)
	count = v
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	rwMtx.Unlock() // 写解锁
}
