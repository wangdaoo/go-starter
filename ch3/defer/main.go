package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("文件创建失败：", err)
		return
	}
	defer file.Close() // 在函数返回时关闭文件

	// 写入数据
	str := "Hello, world!!!"
	n, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println("文件写入失败：", err)
		return
	}
	fmt.Println("写入的字节数：", n)
	fmt.Printf("成功写入 %d 个字节\n", n)

	// 读取数据
	data := make([]byte, n)
	// _, err = os.Open("test.txt")
	_, err = file.Seek(0, 0) // 将文件指针移动到文件开头
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}

	r, err := file.Read(data)
	if err != nil {
		fmt.Println("文件读取失败：", err)
		return
	}
	fmt.Println("读取到的数据：", string(data[:r]))
}
