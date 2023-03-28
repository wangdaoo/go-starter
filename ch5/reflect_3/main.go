package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 10
	var y float64 = 10.0

	// q: reflect.TypeOf()
	// a: reflect.TypeOf() 用于获取变量的类型。
	fmt.Println(reflect.TypeOf(x) == reflect.TypeOf(y)) // 输出 false
}
