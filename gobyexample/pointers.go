package main

import "fmt"

// 值传递
func zeroVal(val int) {
	val = 0
}

// 引用传递
func zeroPtr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	//  &i 语法来取得 i 的内存地址
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)
	fmt.Println("pointer:", &i)
}
