package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ParseFloat 解析浮点数，bitSize: 表示将解析的浮点数类型，返回 float64
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Printf("%T: %f\n", f, f)

	// ParseInt 解析整形数，base:  0 表示自动推断字符串所表示的数字的进制
	i, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%T: %d\n", i, i)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%T: %d\n", u, u)

	// Atoi 是一个基础的 10 进制整型数转换函数
	k, _ := strconv.Atoi("123")
	fmt.Printf("%T: %d\n", k, k)

	// parse error
	_, e := strconv.Atoi("qwe")
	fmt.Println(e)

}
