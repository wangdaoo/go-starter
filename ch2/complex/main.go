package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
)

func main() {

	x := complex(1, 2) // 1+2i
	y := complex(3, 4) // 3+4i

	fmt.Println(x * y) // "(-5+10i)"

	// 类型判断
	fmt.Println(reflect.TypeOf(x))

	// real() 和 imag() 函数分别返回复数的实部和虚部
	fmt.Println(real(x + y))

	fmt.Println(imag(x + y))

	cmplx.Abs(x)              // 数学库中的函数，返回复数的模
	fmt.Println(cmplx.Abs(x)) // "2.23606797749979" (the absolute value)
}
