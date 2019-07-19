package main

import (
	"fmt"
	"sync"
)

type worker struct {
	id   int
	jobs chan int
	done func()
}

func doWork(w worker) {
	for n := range w.jobs {
		fmt.Printf("<Worker:%d>: received %c\n", w.id, n)
		w.done()
	}
}

func createWork(id int, wg *sync.WaitGroup) worker {
	w := worker{id, make(chan int), wg.Done}
	go doWork(w)
	return w
}

func chanDemo() {
	//var c chan int // c == nil

	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i, &wg)
	}

	//wg.Add(20)
	for i, worker := range workers {
		wg.Add(1)
		worker.jobs <- 'a' + i
	}

	for i, worker := range workers {
		wg.Add(1)
		worker.jobs <- 'A' + i
	}
	// 等待任务执行完毕
	wg.Wait()
}

func main() {
	chanDemo()
}
