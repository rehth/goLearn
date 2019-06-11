package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	ch := make(map[int]int)
	// mutex 将同步对 ch 的访问
	mutex := &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += ch[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 确保这个 Go 协程不会在调度中饿死，
				// 在每次操作后明确的使用 runtime.Gosched()进行释放
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for true {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				ch[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("OPS:", opsFinal)

	// 对 ch 使用一个最终的锁
	mutex.Lock()
	fmt.Println("state:", ch)
	mutex.Unlock()
}
