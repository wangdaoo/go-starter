package main

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// range
	for i, v := range a {
		println(i, v)
	}

	println("-------------")

	// range with _
	for _, v := range a {
		println(v)
	}

	println("-------------")

	b := [...]string{"a", "b", "c", "d", "e"}
	for i, v := range b {
		println(i, v)
	}
	// 获取 b 的长度
	println(len(b))

	slice := b[:]

	println(len(slice))

	// slice 插入元素
	slice = append(slice, "f")

	println(cap(slice))

	// 查看 slice 元素
	for i, v := range slice {
		println(i, v)
	}
}
