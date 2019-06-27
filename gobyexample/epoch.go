package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	ns := now.UnixNano()
	ms := ns / 1000000
	fmt.Println("now", now)
	fmt.Println("sec", sec)
	fmt.Println("ms", ms)
	fmt.Println("ns", ns)

	// 将协调世界时起的整数秒或者纳秒转化到相应的时间
	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, ns))
}
