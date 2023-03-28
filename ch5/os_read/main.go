package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// q: 该程序在做什么
// a: 该程序在读取文件。
func main() {
	fileName := "mydir/myfile.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v", fileName, err)
		return
	}
	defer file.Close()

	// q: 为什么要用ioutil.ReadAll
	// a: ioutil.ReadAll会读取文件的所有内容，返回一个字节切片。
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", fileName, err)
		return
	}

	fmt.Printf("Contents of file %s:\n%s", fileName, data)
}
