package main

import (
	"fmt"
	"time"
)

func face(n int) int {
	// 基线条件
	if n <= 0 {
		return 1
	}
	return n + face(n-1)
}

func main() {
	start := time.Now()
	fmt.Println("102400=", face(102400))
	fmt.Println(time.Since(start).Seconds(), "s")
}
