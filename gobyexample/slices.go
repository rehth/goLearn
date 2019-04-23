package main

import "fmt"

func main() {
	// slice 的类型仅由它所包含的元素决定
	s := make([]string, 3)
	fmt.Println("make", s)

	s[0] = "a"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// append返回的新的 slice 值
	s = append(s, "d", "e")
	fmt.Println("append:", s)

	// Slice 也可以被 copy
	c := make([]string, len(s))
	copy(c, s[:3])
	fmt.Println("copy:", c)
	c[1] = "b"
	fmt.Println("set:", c, s)

	// Slice 支持通过 slice[low:high] 语法进行“切片”操作
	fmt.Println("sl1:", s[2:4])
	fmt.Println("sl2:", s[:4])

	// slice 支持字面值初始化
	fmt.Println("init:", []string{"a", "c", "F"})

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
