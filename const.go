package main

import "fmt"

// 常量值必须是编译期可确定的数字、字符串、布尔值。
const (
	A = iota  			// 0 支持多个 iota，它们各自增长
	B = iota 			// 1
	C					// 2 可省略后续的表达式，表示同上一行表达式一致。
	D              		// 3
	E    = "k"        	// "K"
	F    = iota       	// 5 iota 自增被打断，须显式恢复。
	G                 	// 6

)

func main() {
	const H = "H"  // 未使用的常量不会产生错误，支持类型推断、并行赋值、常量组、枚举等
	fmt.Println(A, B, C, D, E, F, G)
}
