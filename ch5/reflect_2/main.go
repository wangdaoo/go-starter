package main

import (
	"fmt"
	"reflect"
)

// q: Calculator 类型
// a: Calculator 类型是一个结构体类型，它有一个 Add 方法，该方法接收两个 int 类型的参数，返回一个 int 类型的结果。
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func main() {
	c := Calculator{}
	value := reflect.ValueOf(c)

	// q: MethodByName()
	// a: MethodByName() 用于获取结构体类型的方法，这里是获取 Add 方法。
	method := value.MethodByName("Add")
	if method.IsValid() {
		// q: Call()
		// a: Call() 用于调用反射值的方法，这里是调用 Add 方法。
		result := method.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})
		// q: Int()
		// a: Int() 用于获取反射值的 int 类型的值，这里是获取 Add 方法的返回值。
		fmt.Println(result[0].Int()) // 输出 3
	}
}
