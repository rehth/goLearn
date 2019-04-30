package main

import "fmt"

func main() {
	// 默认通道是 无缓冲 的
	// make 了一个通道，最多允许缓存 2 个值
	ch := make(chan string, 2)

	ch <- "buff"
	ch <- "chan"

	fmt.Println(1, <-ch)
	fmt.Println(2, <-ch)
}
