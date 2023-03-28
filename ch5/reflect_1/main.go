package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 20}
	value := reflect.ValueOf(&p).Elem() // 获取结构体指针的反射值

	name := value.FieldByName("Name")
	// q: IsValid()
	// a: IsValid() 用于判断反射值是否有效，如果反射值是无效的，那么后续的操作都是不安全的。
	if name.IsValid() {
		// q: CanSet()
		// a: CanSet() 用于判断反射值是否可修改，如果反射值是不可修改的，那么后续的修改操作都是不安全的。
		if name.CanSet() {
			// q: SetString()
			// a: SetString() 用于修改反射值的值，这里是修改 Name 字段的值。
			name.SetString("Bob") // 修改 Name 字段的值
		}
	}

	age := value.FieldByName("Age")
	if age.IsValid() && age.CanSet() {
		age.SetInt(25) // 修改 Age 字段的值
	}

	fmt.Println(p) // 输出 {Bob 25}
}
