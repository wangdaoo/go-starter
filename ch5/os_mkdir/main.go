package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// q: 该程序在做什么
// a: 该程序在创建目录和写入文件。
func main() {
	dirName := "mydir"
	fileName := "myfile.txt"

	// 创建目录
	// q: 为什么要用0755
	// a: 0755是一个八进制数，表示的是二进制的111101101，也就是rwxr-xr-x，也就是所有者有读写执行权限，组用户有读执行权限，其他用户有读执行权限。
	if err := os.Mkdir(dirName, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v", dirName, err)
		return
	}

	// 写入文件
	data := []byte("hello world\n")
	// q: 为什么要用0644
	// a: 0644是一个八进制数，表示的是二进制的110100100，也就是rw-r--r--，也就是所有者有读写权限，组用户有读权限，其他用户有读权限。
	if err := ioutil.WriteFile(dirName+"/"+fileName, data, 0644); err != nil {
		fmt.Printf("Error writing file %s/%s: %v", dirName, fileName, err)
		return
	}

	fmt.Printf("Successfully created directory %s and wrote file %s/%s", dirName, dirName, fileName)
}
