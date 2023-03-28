package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	wg    sync.WaitGroup
	mutex sync.Mutex // 互斥锁
)

func main() {
	wg.Add(2)

	go incCount()
	go incCount()

	wg.Wait()
	// 不加锁的情况下，count 的值可能不是 4
	fmt.Println("Final count:", count) // Final count: 4
}

func incCount() {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// q: 为什么这里要加锁？
		// a: 因为 count 变量是共享的，如果不加锁，可能会出现多个 Goroutine 同时读写 count 变量的情况，导致 count 变量的值不正确。
		mutex.Lock() // 加锁
		{
			value := count
			time.Sleep(time.Millisecond * 1)
			value++
			count = value
		}
		// q: 为什么这里要解锁？
		// a: 因为加锁和解锁是成对出现的，如果不解锁，那么其他 Goroutine 就无法获取到锁，从而无法执行。
		mutex.Unlock() // 解锁
	}
}
