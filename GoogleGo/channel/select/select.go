package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c <-chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("<worker %d>: received %d\n", id, v)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var c1, c2 chan int // c1 and c2 == nil
	c1, c2 := generator(), generator()
	w := createWorker(0)

	timer := time.After(10 * time.Second)
	tick := time.Tick(time.Second * 2)

	var values []int
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
			//fmt.Println("Received from C1:", n)
		case n := <-c2:
			values = append(values, n)
			//fmt.Println("Received from C2:", n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Received timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-timer:
			fmt.Println("bye bye")
			return

			//default:   非堵塞模式
			//	fmt.Println("No value received")
		}
	}
}
