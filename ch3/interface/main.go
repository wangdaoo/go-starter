package main

import (
	"fmt"
	"math"
)

// 定义一个接口
type geometry interface {
	area() float64
	perim() float64
}

// 定义一个矩形结构体
type rect struct {
	width, height float64
}

// 实现矩形的面积和周长方法
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 定义一个圆形结构体
type circle struct {
	radius float64
}

// 实现圆形的面积和周长方法
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 定义一个计算几何图形面积和周长的函数
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// 假设有一个服务需要从不同的数据源中获取数据，数据源可以是 MySQL 数据库、Redis 缓存或者是文件系统。
// 为了使代码更加灵活和可扩展，可以使用 interface 来定义一个 DataSource 接口，并让每个数据源实现这个接口。
// 然后在代码中使用这个接口类型的变量来获取数据，而不需要关心具体使用了哪个数据源。
type DataSource interface {
	GetData() string
}

type MySQL struct {
	// 定义 MySQL 数据源的属性
	types string
	// 定义一个获取当前时间戳的方法
	GetTimestamp func() int64
}

func (m MySQL) GetData() string {
	return "MySQL ->" + m.types + " -> " + fmt.Sprintf("%d", m.GetTimestamp())
}

type Redis struct {
	// 定义一个两数相加的方法
	Add func(a, b int)
}

func (r Redis) GetData() string {
	r.Add(1, 2)
	return "Redis"
}

type File struct{}

func (f File) GetData() string {
	return "File"
}

func GetDataFromSource(source DataSource) string {
	return source.GetData()
}

func getTime() int64 {
	return 1234567890
}

func handleAdd(a, b int) {
	fmt.Println("Add: a + b =", a+b)
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 通过接口调用计算几何图形面积和周长的函数
	measure(r)
	measure(c)

	m := MySQL{types: "mysql", GetTimestamp: getTime}
	rs := Redis{Add: handleAdd}

	// 通过接口调用获取数据的函数
	fmt.Println(GetDataFromSource(m))
	fmt.Println(GetDataFromSource(rs))
	fmt.Println(GetDataFromSource(File{}))
}
