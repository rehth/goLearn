package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// range 迭代从 queue 中得到的每个值
	// 若 通道 没有close，则继续阻塞执行
	for elem := range queue {
		fmt.Println(elem)
	}
}
