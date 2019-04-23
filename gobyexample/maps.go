package main

import "fmt"

func main() {
	m := make(map[string]int, 10)

	//  make[key] = val 语法来设置键值对
	m["k1"] = 7
	m["k2"] = 15
	fmt.Println("map:", m)

	// name[key] 来获取一个键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// len 时，返回的是键值对数目
	fmt.Println("len:", len(m))

	//  delete 可以从一个 map 中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	//  map 中取值时，可选的第二返回值指示这个键是在这个 map 中
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 申明和初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
