package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ------------------------------- 读写锁 --------------------------------
// 1. 在读锁定期间，其他 goroutine 可以同时读取共享变量，但不能写入共享变量。
// 2. 在写锁定期间，其他 goroutine 不能同时读取或写入共享变量。
// 3. 读写锁在写锁定期间是排他的，这意味着其他 goroutine 不能同时持有读锁或写锁。
// ------------------------------- 读写锁 --------------------------------

// 共享变量
var data = make(map[string]string)

// 读写锁
var rwLock sync.RWMutex

// 读取共享变量
func readData(key string) string {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return data[key]
}

// 写入共享变量
func writeData(key, value string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	data[key] = value
}

// 随机生成一个字符串
func randString(n int) string {
	// q: rune 是什么？
	// a: rune 是 int32 的别名，用于表示 Unicode 码点。
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		// q: rand.Intn() 是什么？
		// a: rand.Intn() 返回一个随机的整数 n，0 <= n < len(letters)。
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 读取共享变量的 goroutine
func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := readData(key)
		fmt.Printf("reader%d read %s: %s\n", id, key, value)
		time.Sleep(100 * time.Millisecond)
	}
}

// 写入共享变量的 goroutine
func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := randString(10)
		writeData(key, value)
		fmt.Printf("writer%d write %s: %s\n", id, key, value)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i, &wg)
	}
	wg.Wait()
}
