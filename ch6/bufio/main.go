package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个新的读取器，从标准输入读取数据
	// q: os.Stdin
	// a: os.Stdin 是一个文件指针，指向标准输入
	reader := bufio.NewReader(os.Stdin)

	// 读取一行数据
	fmt.Print("请输入一行文字：")
	// q: reader.ReadString('\n')
	// a: 从 reader 中读取数据，直到遇到 '\n' 为止
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("读取到的数据：", line)

	// 读取一个字节
	fmt.Print("请输入一个字符：")
	// q: reader.ReadByte()
	// a: 从 reader 中读取一个字节
	char, err := reader.ReadByte()
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("读取到的数据：", string(char))

	// 缓冲区写入数据
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprint(writer, "Hello, ")
	fmt.Fprint(writer, "World!\n")
	// q: writer.Flush()
	// a: 将缓冲区中的数据写入到 os.Stdout 中
	writer.Flush()
}
