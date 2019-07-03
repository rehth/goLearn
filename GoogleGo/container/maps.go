package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	start, maxLength := 0, 0
	lastOccurred := make(map[rune]int)

	for ind, c := range s {
		if last, ok := lastOccurred[c]; ok && last >= start {
			start = last + 1
		}
		if ind-start+1 > maxLength {
			maxLength = ind - start + 1
		}
		lastOccurred[c] = ind
	}
	return maxLength
}

func main() {
	// 除了 slice, map, function 的内建类型都可作为key
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m2 := make(map[string]string) // m2 is empty map

	var m3 map[string]float64 // m3 is nil

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	fmt.Println("Getting values")
	a := m1["a"]
	d, ok := m1["d"]
	fmt.Println(a, d, ok)

	fmt.Println("Deleting values")
	delete(m1, "c")
	fmt.Println(m1)

	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
	fmt.Println(lengthOfNonRepeatingSubStr("abc"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcabc"))
	fmt.Println(lengthOfNonRepeatingSubStr("这是在学习Golang"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))
}
