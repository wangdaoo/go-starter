package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{Name: "Go"}

	// fmt.Printf("site: %v\n", site) // Output: site: {Go}
	fmt.Printf("site: %#v\n", site) // Output: site: main.WebSite{Name:"Go"}

	fmt.Printf("site: %T\n", site) // Output: site: main.WebSite

	// 8进制输出
	fmt.Printf("site: %o\n", 8) // Output: 10

	// 16进制输出
	fmt.Printf("site: %x\n", 16) // Output: 10

	// 指针输出
	fmt.Printf("site: %p\n", &site) // Output: 0xc00000a0a0
}
