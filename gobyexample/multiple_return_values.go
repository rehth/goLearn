package main

import "fmt"

func values() (int, string) {
	return 10, "11"
}

func main() {
	// 通过多赋值 操作来使用这两个不同的返回值
	a, b := values()
	fmt.Printf("%d, %T\n", a, a)
	fmt.Printf("%s, %T\n", b, b)

	// 可以使用空白定义符 _
	_, c := values()
	fmt.Println(c)
}
