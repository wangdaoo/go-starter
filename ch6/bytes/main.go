package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 将字符串转换为字节切片
	b := []byte("Hello, world!")

	// 将字节切片转换为字符串
	fmt.Println(string(b))

	// 将字节切片中的所有字符转换为大写
	fmt.Println(bytes.ToUpper(b))

	// 将字节切片中的所有字符转换为小写
	fmt.Println(bytes.ToLower(b))

	// 判断两个字节切片是否相等
	fmt.Println(bytes.Equal([]byte("hello"), []byte("world")))
	fmt.Println(bytes.Equal([]byte("hello"), []byte("hello")))

	// 判断字节切片是否包含另一个字节切片
	fmt.Println(bytes.Contains(b, []byte("world")))
	fmt.Println(bytes.Contains(b, []byte("Mars")))

	// 将字节切片中的所有字符替换为另一个字符
	fmt.Println(bytes.Replace(b, []byte("o"), []byte("0"), -1))

	// 将字节切片中的所有字符替换为另一个字符，只替换一次
	fmt.Println(bytes.Replace(b, []byte("o"), []byte("0"), 1))

	// 将字节切片中的所有字符替换为另一个字节切片
	fmt.Println(bytes.ReplaceAll(b, []byte("o"), []byte("0")))

	// 将字节切片中的所有字符拼接成一个字符串
	fmt.Println(string(bytes.Join([][]byte{[]byte("Hello"), []byte("world")}, []byte(", "))))

	// 将字节切片中的所有字符拼接到另一个字节切片中
	fmt.Println(bytes.Join([][]byte{[]byte("Hello"), []byte("world")}, []byte(", ")))

}
