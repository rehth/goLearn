package main

import "fmt"

// 使用任意数目的 int 作为参数
func sum(ns ...int) {
	fmt.Print(ns, " ")
	total := 0
	for _, n := range ns {
		total += n

	}
	fmt.Println(total)
}

func main() {
	// 常规的调用方式
	sum(1, 2)
	sum(1, 2, 3)
	num := []int{1, 2, 3, 4, 5}
	// 调用 func(slice...)
	sum(num...)
}
