package main

import (
	"fmt"
	"math"
)

// 定义一个结构体
type Vertex struct {
	X, Y float64
}

// 定义一个方法，接收者为 Vertex 类型
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 定义一个方法，接收者为 *Vertex 类型
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 定义一个方法，接收者为 *Vertex 类型
func (v *Vertex) Normalize() {
	abs := v.Abs()
	v.X = v.X / abs
	v.Y = v.Y / abs
}

type person struct {
	name string
	age  int
}

// 实现继承, person 类型实现了 GetName() 方法
type human struct {
	person
}

func (p person) GetName() string {
	return p.name
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("原始值：", v)            // {3 4}
	fmt.Println("Abs() 方法：", v.Abs()) // 5

	v.Scale(10)
	fmt.Println("Scale() 方法：", v) // {30 40}

	v.Normalize()
	fmt.Println("Normalize() 方法：", v) // {0.6 0.8}

	p := person{name: "Bob", age: 20}
	fmt.Println("原始值：", p) // {Bob 20}

	fmt.Println("GetName() 方法：", p.GetName()) // Bob

	h := human{}
	h.name = "Alice"
	h.age = 18
	fmt.Println("原始值：", h)
	fmt.Println("GetName() 方法：", h.GetName())
}
