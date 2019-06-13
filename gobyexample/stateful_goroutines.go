package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var ops int64                 // 记录操作数
	reads := make(chan *readOp)   // 记录其他 goroutine 发布的 read 请求
	writes := make(chan *writeOp) // 记录其他 goroutine 发布的 write 请求

	go func() {
		// 读取并处理 read/write 请求
		// state 将被一个单独的 Go 协程拥有，保证数据在并行读取时不会混乱
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true

			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			// 发送 read 请求 by 100 goroutine
			for {
				read := &readOp{key: rand.Intn(5), resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			// 发送 write 请求 by 10 goroutine
			for {
				write := &writeOp{key: rand.Intn(5), val: rand.Intn(100), resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	// 让所有go协程运行1s
	time.Sleep(time.Second)
	// 查看操作次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops", opsFinal)
}
