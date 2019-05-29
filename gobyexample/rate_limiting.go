package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	quit := make(chan bool)
	limiter := time.NewTicker(time.Millisecond * 200)

	for req := range requests {
		<-limiter.C
		fmt.Println("limit request,", req, time.Now())
	}

	bursty := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		bursty <- time.Now()
	}

	go func() {
	LOOP:
		for {
			select {
			case t := <-limiter.C:
				//fmt.Println("token")
				bursty <- t
			case <-quit:
				// 退出 for-select 语句
				break LOOP
			}
		}
		fmt.Println("goroutine stopped")
	}()

	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-bursty
		fmt.Println("request", req, time.Now())
	}
	limiter.Stop()
	quit <- true
	time.Sleep(time.Second)

}
