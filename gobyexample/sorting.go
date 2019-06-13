package main

import (
	"fmt"
	"sort"
)

func main() {
	// str 排序，排序是原地更新的
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("SortStr:", strs)

	// int 排序
	ints := []int{1, 4, 2, 6, 5}
	sort.Ints(ints)
	fmt.Println("SortInt:", ints)

	// 检查一个序列是不是已经是排好序的。
	fmt.Println("Sorted:", sort.IntsAreSorted(ints))
}
