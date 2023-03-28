package main

import (
	"fmt"
	"math"
)

// 1. 基本函数定义
func add(a, b int) int {
	return a + b
}

// 2. 多返回值函数
func swap(a, b string) (string, string) {
	return b, a
}

// 3. 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9 // 17 * 4 / 9 = 7
	y = sum - x     // 17 - 7 = 10
	return
}

// 4. 变参函数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 5. 闭包函数
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 6. 递归函数
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// 7. 函数作为参数
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 8. 函数作为返回值
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 9. defer 延迟执行
func deferDemo() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// 10. panic 和 recover
func panicDemo() {
	defer func() {
		str := recover() // recover 用于捕获 panic，防止程序崩溃。
		fmt.Println(str)
	}()
	panic("crash")
}

// 11. 匿名函数
func anonymousDemo() {
	func() {
		fmt.Println("anonymous function")
	}()
}

// 12. 方法
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) // 3*3 + 4*4 = 25 -> 5
}

// 13. 指针接收者的方法
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 14. 接口
type Abser interface {
	Abs() float64
}

// 15. 实现接口
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func interfaceDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f                // a MyFloat 实现了 Abser
	fmt.Println(a.Abs()) // 1.4142135623730951，调用的是 MyFloat 的 Abs 方法

	a = &v               // a *Vertex 实现了 Abser
	fmt.Println(a.Abs()) // 5，调用的是 Vertex 的 Abs 方法
}

func main() {
	// 1. 基本函数定义
	fmt.Println(add(42, 13)) // 55

	// 2. 多返回值函数
	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	// 3. 命名返回值
	fmt.Println(split(17)) // 7 10

	// 4. 变参函数
	sum(1, 2)    // [1 2] 3
	sum(1, 2, 3) // [1 2 3] 6

	nums := []int{1, 2, 3, 4} // [1 2 3 4] 10
	sum(nums...)

	// 5. 闭包函数
	nextInt := intSeq()    // -> 1
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()    // -> 1
	fmt.Println(newInts()) // 1

	// 6. 递归函数
	fmt.Println(fact(7)) // 5040

	// 7. 函数作为参数
	fmt.Println(compute(math.Pow)) // 81。 来源：math.Pow(3, 4)，3 的 4 次方

	// 8. 函数作为返回值
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // 0 1 3 6 10 15 21 28 36 45
			neg(-2*i), // 0 -2 -6 -12 -20 -30 -42 -56 -72 -90
		)
	}
	// 函数作为返回值
	// adder 函数返回一个函数，这个函数的参数是 int 类型，返回值也是 int 类型。
	// 在 main 函数中，我们将 adder 函数赋值给 pos 和 neg 两个变量，这两个变量都是函数类型。
	// 在 for 循环中，我们分别调用 pos 和 neg 函数，每次调用都会将参数累加到 sum 变量中，然后返回 sum 变量的值。
	// 因此，pos(i) 的返回值是 0、1、3、6、10、15、21、28、36、45，neg(-2*i)
	// 的返回值是 0、-2、-6、-12、-20、-30、-42、-56、-72、-90。

	// 9. defer 延迟执行
	deferDemo() // hello world。defer 语句会将函数推迟到外层函数返回之后执行。

	// 10. panic 和 recover
	panicDemo() // crash。panic 会中断原有的控制流程，进入一个 panic 状态，直到当前 goroutine 中所有的延迟函数执行完毕。

	// 11. 匿名函数
	anonymousDemo() // anonymous function。 匿名函数没有函数名，但是可以在函数体内部使用。

	// 12. 方法
	v := Vertex{3, 4}

	fmt.Println(v.Abs()) // 5, 来源：math.Sqrt(3*3 + 4*4). math.Sqrt() 函数是 Go 语言的标准库函数, 用于计算平方根

	// 13. 指针接收者的方法
	v.Scale(10)
	fmt.Println(v.Abs())
	// 指针接收者的方法
	// 在 main 函数中，我们定义了一个 Vertex 类型的变量 v，
	// 然后调用 v 的 Scale 方法，将 v 的 X 和 Y 坐标都乘以 10。
	// 最后再次调用 v 的 Abs 方法，计算 v 的模长。由于 Scale 方法的接收者是指针类型，
	// 因此 v 的 X 和 Y 坐标都被修改了，Abs 方法计算的是修改后的坐标，结果为 50。
	// 勾股定理：a² + b² = c²

	// 14. 接口
	interfaceDemo()
}
