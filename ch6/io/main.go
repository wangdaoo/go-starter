package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func copy() {
	// q: strings.NewReader() 是什么意思？
	// a: strings.NewReader() 返回一个指向字符串的指针
	r := strings.NewReader("Hello, Reader!\n")

	// q: io.Copy() 是什么意思？
	// a: io.Copy() 将 r 中的内容复制到 w 中
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func copyBuffer() {
	r := strings.NewReader("Hello, Reader!\n")
	buf := make([]byte, 8)
	// q: io.CopyBuffer() 是什么意思？
	// a: io.CopyBuffer() 将 r 中的内容复制到 w 中，使用 buf 作为缓冲区
	// q: 为什么要使用缓冲区？
	// a: 为了提高效率, 减少系统调用, 降低内存分配, 降低 GC 压力. 适用于大量数据的复制
	_, err := io.CopyBuffer(os.Stdout, r, buf)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	copy()

	copyBuffer()
}
