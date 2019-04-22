package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	// 数值型常量是没有确定的类型的，直到他们需要一个类型时
	const n = 50000000000

	// const 可以出现在 var 可以出现的任意地方
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
