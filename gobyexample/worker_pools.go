package main

import (
	"fmt"
	"time"
)

// 模仿一个耗时的任务
func work(id int, jobs <-chan int, res chan<- int) {
	// 从任务队列中取任务
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		res <- j * 2
	}
}

func main() {
	workNum := 12  // 任务数
	threadNum := 3 // 并发数
	jobs := make(chan int, 100)
	res := make(chan int, 100)

	// 启动了 3 个 worker，初始是阻塞的
	for w := 1; w <= threadNum; w++ {
		go work(w, jobs, res)
	}
	// 发送 12 个 jobs
	for j := 1; j <= workNum; j++ {
		jobs <- j
	}
	//  close 这些通道来表示这些就是所有的任务了
	close(jobs)

	// 获取任务结果
	for a := 1; a <= workNum; a++ {
		<-res
	}
	fmt.Println("ending...")
}
