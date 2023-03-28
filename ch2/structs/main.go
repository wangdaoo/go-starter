package main

import (
	"fmt"
	"time"
)

// 定义一个结构体
type person struct {
	name string
	age  int
	Dob  time.Time
	t    string
}

func showPerson(p person) {
	fmt.Println("-----showPerson-----")
	fmt.Println(p)
	fmt.Println("-----showPerson-----")
}

func main() {
	// 初始化结构体
	p1 := person{
		name: "Alice",
		age:  20,
		Dob:  time.Now(),
		t:    time.Now().Format("2006-01-02 15:04:05"),
	}
	fmt.Println(p1)

	// 访问结构体字段
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	showPerson(p1)

	// 结构体指针
	p2 := &person{name: "Bob", age: 30}
	fmt.Println(p2)
	// p2.name === (*p2).name，Go 语言会自动帮我们解引用，所以可以直接写 p2.name
	fmt.Println((*p2).name)
	fmt.Printf("p2.name: %v\n", p2.name)
	fmt.Println(p2.age)

	// 结构体嵌套
	type address struct {
		city  string
		state string
	}
	type contact struct {
		email   string
		phone   string
		address address
	}
	c1 := contact{
		email: "alice@example.com",
		phone: "1234567890",
		address: address{
			city:  "New York",
			state: "NY",
		},
	}
	fmt.Println(c1)
	fmt.Println(c1.email)
	fmt.Println(c1.phone)
	fmt.Println(c1.address.city)
	fmt.Println(c1.address.state)

	// 匿名结构体
	p3 := struct {
		name string
		age  int
	}{
		name: "Charlie",
		age:  40,
	}
	fmt.Println(p3)
	fmt.Println(p3.name)
	fmt.Println(p3.age)
}
