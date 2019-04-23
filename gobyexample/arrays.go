package main

import "fmt"

func main() {
	// 元数组类型是由元素的类型和数组长度组成
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// 支持进行初始化赋值
	b := [5]int{1: 4, 4: 1, 0: 5}
	fmt.Println("init:", b)

	// 支持多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
