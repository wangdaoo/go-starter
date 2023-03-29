package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// 使用字节数组创建一个新的 Reader
	data := []byte("Hello, World!")
	r := bytes.NewReader(data)

	// 读取单个字节
	b, err := r.ReadByte()
	if err != nil {
		fmt.Println("ReadByte error:", err)
		return
	}
	fmt.Printf("ReadByte: %c\n", b)

	// 读取多个字节
	buf := make([]byte, 5)
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Printf("Read: %s\n", buf[:n])

	// 重置 Reader，使其重新从头开始读取
	r.Reset(data)

	// 查看当前位置
	pos, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("Seek error:", err)
		return
	}
	fmt.Println("Current position:", pos)

	// 读取 Reader 中剩余的字节数
	n = r.Len()
	fmt.Println("Remaining bytes:", n)
}
