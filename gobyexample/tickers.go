package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器 是当你想要在固定的时间间隔重复执行准备的
	ticker := time.NewTicker(time.Millisecond * 500)
	quit := make(chan bool)

	go func() {
		// 使用 label 来组织 for-select 语句
	LOOP:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("ticker at:", t)
			case <-quit:
				// 退出 for-select 语句
				break LOOP
			}
		}

		fmt.Println("goroutine stopped")
	}()

	time.Sleep(time.Millisecond * 1600)
	// 打点器可以和定时器一样被停止
	ticker.Stop()
	quit <- true
	fmt.Println("ticker stopped")
	time.Sleep(time.Second)
}
