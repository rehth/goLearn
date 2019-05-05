package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func random(n int) int {
	return rand.Intn(n)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	rand.Seed(time.Now().UnixNano())
	go func() {
		n := random(4)
		time.Sleep(time.Second * time.Duration(n))
		c1 <- "one" + strconv.Itoa(n)
	}()

	go func() {
		m := random(4)
		time.Sleep(time.Second * time.Duration(m))
		c2 <- "two" + strconv.Itoa(m)
	}()

	for i := 0; i < 2; i++ {
		// Go 的通道选择器 让你可以同时等待多个通道操作
		select {
		case msg1 := <-c1:
			fmt.Println("c1", msg1)
		case msg2 := <-c2:
			fmt.Println("c2", msg2)
		}

	}
}
