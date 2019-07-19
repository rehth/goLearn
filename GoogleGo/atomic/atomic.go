package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	mux   sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	func() {
		a.mux.Lock()
		defer a.mux.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
