package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	fmt.Print("请输入：")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")

	}

	switch time.Now().Weekday() {
	// case 语句中，可以使用逗号来分隔多个表达式
	case time.Saturday, time.Sunday:
		fmt.Println("愉快的周末")
	default:
		fmt.Println("工作日，你懂得。。。")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("这是上午")
	default:
		fmt.Println("下午")

	}

}
