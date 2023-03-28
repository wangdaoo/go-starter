package main

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func (c *Counter) Incr() {
	// q: atomic.AddInt64()
	// a: atomic.AddInt64() 用于原子地增加一个 int64 类型的值。
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) Value() int64 {
	// q: atomic.LoadInt64()
	// a: atomic.LoadInt64() 用于原子地获取一个 int64 类型的值。
	return atomic.LoadInt64(&c.value)
}

func main() {
	var counter Counter

	for i := 0; i < 100; i++ {
		go func() {
			counter.Incr()
		}()
	}

	// 等待所有 goroutine 执行完毕
	for counter.Value() < 100 {
	}

	fmt.Println(counter.Value()) // 输出 100
}
