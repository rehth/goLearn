package main

import (
	"fmt"
	"strings"
)

//  strings 切片的组合函数示例: Index, Include, Any, All, Filter, Map
func Index(vs []string, t string) int {
	for ind, v := range vs {
		if v == t {
			return ind
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	fer := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			fer = append(fer, v)
		}
	}
	return fer
}

func Map(vs []string, f func(string) string) []string {
	mp := make([]string, len(vs))
	for i, v := range vs {
		mp[i] = f(v)
	}
	return mp
}

func main() {
	fruits := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(fruits, "pear"))

	fmt.Println(Include(fruits, "grape"))

	fmt.Println(Any(fruits, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))

	fmt.Println(All(fruits, func(i string) bool {
		return strings.HasPrefix(i, "p")
	}))

	fmt.Println(Filter(fruits, func(i string) bool {
		return strings.Contains(i, "e")
	}))

	fmt.Println(Map(fruits, strings.ToUpper))
}
