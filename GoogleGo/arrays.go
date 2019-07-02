package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	// 数组是值类型
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 5, 7, 4}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	fmt.Println("printArray: arr1")
	printArray(&arr1)
	fmt.Println("printArray: arr3")
	printArray(&arr3)
	fmt.Println(arr1, arr3)
}
