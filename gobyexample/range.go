package main

import "fmt"

func main() {
	ns := []int{1, 2, 3}
	var sum int
	for _, n := range ns {
		sum += n
	}
	fmt.Println("sum:", sum)

	// range 在数组和 slice 中都同样提供每个项的索引和值
	for i, n := range ns {
		fmt.Println("ind:", i, "val:", n)
	}

	// range 在 map 中迭代键值对
	kvs := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range kvs {
		fmt.Printf("%s --> %d\n", k, v)
	}

	// range 在字符串中迭代 unicode
	for i, c := range "golang" {
		fmt.Printf("str:%d %c\n", i, c)
	}
}
