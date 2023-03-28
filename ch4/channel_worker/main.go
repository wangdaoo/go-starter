package main

// Test channel with worker goroutine

import (
	"fmt"
	"time"
)

// q: results 在做什么事情？
// a: results channel 用于接收 worker goroutine 的执行结果
// q: jobs 在做什么事情？
// a: jobs channel 用于接收 main goroutine 的任务
// q: 为什么要用 channel 来传递任务？
// a: 通过 channel 来传递任务，可以让 main goroutine 和 worker goroutine 之间解耦，main goroutine 只需要把任务发送到 jobs channel，不需要关心任务是如何被执行的，而 worker goroutine 只需要从 jobs channel 中取任务，不需要关心任务是从哪里来的
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

// 1. 创建 jobs 和 results 两个 channel
// 2. 启动3个 worker goroutine，每个 goroutine 都会从 jobs channel 中取任务，执行完后将结果写入 results channel
// 3. 发送5个任务到 jobs channel
// 4. 接收5个结果
// 5. 关闭 jobs channel
// 6. 打印结果
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个 worker goroutine
	// q: 此时 worker 循环是否无序执行？
	// a: worker 循环是无序执行的，因为 worker goroutine 是并发执行的
	// q: 为什么要用 for 循环来启动 worker goroutine？
	// a: 为了让 main goroutine 等待所有 worker goroutine 执行完毕，如果不用 for 循环，main goroutine 可能会在 worker goroutine 执行完毕之前退出
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送5个任务到 jobs channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 接收5个结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
