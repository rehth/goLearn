package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址
				atomic.AddUint64(&ops, 1)
				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	// 通过 LoadUint64 将当前值的拷贝提取到 opsFinal中。
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops", ops, opsFinal)
}
