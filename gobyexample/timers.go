package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器表示在未来某一时刻的独立事件
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	// timer 可以被停止
	// 使用timer.Reset()代替timer.Stop()来停止定时器
	if stop2 := timer2.Reset(0); stop2 {
		fmt.Println("Timer2 stopped")
	}
	time.Sleep(time.Second)
}
