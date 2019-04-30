package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//  Go 协程 在执行上来说是轻量级的线程
	//  Go 运行时是以异步的方式运行协程的
	f("start")

	//  Go 协程将会并行的执行这个函数调用
	go f("goroutine")

	// 可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scan(&input)
	fmt.Println(input, "end")
}
