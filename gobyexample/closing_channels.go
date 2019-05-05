package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 如果 jobs 已经关闭了，并且通道中所有的值都已经接收完毕，
			// 那么 more 的值将是 false
			if j, more := <-jobs; more {
				fmt.Println("一个工作：", j)
			} else {
				fmt.Println("没有多余的工作了")
				done <- true
				return
			}

		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("添加了一份工作：", j)
	}
	// 关闭 一个通道意味着不能再向这个通道发送值了
	close(jobs)
	fmt.Println("不在添加新的工作")
	// 使用 通道同步方法等待任务结束
	<-done

}
