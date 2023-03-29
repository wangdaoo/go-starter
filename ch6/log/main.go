package main

import (
	"log"
	"os"
)

func main() {
	// 创建一个文件，用于记录日志
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("无法创建日志文件：", err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 设置日志前缀
	log.SetPrefix("【Info】")

	// 设置日志输出格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 输出日志
	log.Println("这是一条普通日志。")
	log.Fatalln("这是一条错误日志。")
	log.Println("这条日志不会输出。")
}
