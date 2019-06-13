package main

import (
	"fmt"
	"sort"
)

//  sort.Interface 的 Len，Less和 Swap 方法
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	// 想按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println("sort:", fruits)
}
