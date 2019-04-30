package main

import (
	"fmt"
	"time"
)

// 通道将被用于通知其他 Go 协程这个函数已经工作完毕
func worker(ch chan bool) {
	fmt.Println("worker is working...")
	time.Sleep(2 * time.Second)
	fmt.Println("pong")

	ch <- true
}

func main() {
	// 使用通道来同步 Go 协程间的执行状态
	msgBox := make(chan bool, 1)
	go worker(msgBox)

	// 接收到通道中 worker 发出的通知前一直阻塞
	fmt.Println(<-msgBox)
}
