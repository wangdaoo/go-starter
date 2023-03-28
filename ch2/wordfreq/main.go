package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 1. 启动程序，输入一些单词，然后按Ctrl+D结束输入
 2. 程序会统计每个单词出现的次数
*/

func main() {
	wordFreq := make(map[string]int)    // 创建 Map, 用于存储单词和出现的次数
	input := bufio.NewScanner(os.Stdin) // 创建 Scanner 扫描器, 从标准输入读取数据
	input.Split(bufio.ScanWords)        // 设置扫描器的分割函数为 ScanWords, 按单词分割
	for input.Scan() {                  // 循环扫描输入的单词
		wordFreq[input.Text()]++ // 统计单词出现的次数
	}
	if err := input.Err(); err != nil { // 检查扫描器的错误
		fmt.Fprintf(os.Stderr, "wordFreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\tfrequency\n")    // 打印表头
	for word, freq := range wordFreq { // 打印单词和出现的次数
		fmt.Printf("%s\t%d\n", word, freq) // 输出格式为: 单词 出现的次数
	}
}
