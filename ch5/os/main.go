package main

import (
	"fmt"
	"os"
)

// q: 该程序的输出是什么？
// a: 该程序的输出是当前进程的环境变量。
func main() {
	envVars := os.Environ()

	fmt.Println("Environment variables:")
	for _, envVar := range envVars {
		fmt.Println(envVar)
	}
}
