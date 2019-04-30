package main

import "fmt"

func main() {
	// channel 连接多个 Go 协程的管道
	// 使用 make(chan val-type) 创建一个新的通道
	ch := make(chan string)

	go func() {
		// 使用 channel <- 语法 发送（阻塞） 一个新的值到通道中
		ch <- "ping"
	}()
	// 使用 <-channel 语法从通道中 接收（阻塞）  一个值
	msg := <-ch
	fmt.Println("msg:", msg)
}
