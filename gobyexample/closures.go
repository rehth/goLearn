package main

import "fmt"

func intSeq() func() int {
	var i int
	// 返回的函数使用闭包的方式 隐藏 变量 i
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	// 闭包：函数及所需环境变量所组成的整体
	fmt.Println("next", nextInt())
	fmt.Println("next", nextInt())
	fmt.Println("next", nextInt())

	newInt := intSeq()
	fmt.Println("new", newInt())
}
