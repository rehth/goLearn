package main

import (
	"fmt"
	"time"
)

func main() {
	//var s = [10]int{}
	for i := 0; i < 10; i++ {
		// 非抢占式 多任务处理，由协程主动交出控制权
		go func(i int) { // race condition
			for true {
				// I/O 操作 可能有 goroutine 的切换
				fmt.Println("Hello from goroutine:", i)

				//s[i] ++
				//runtime.Gosched() // 主动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Second)
	//fmt.Println(s)
}
