package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// 初始化一个结构体元素时指定字段名字，省略的字段将被初始化为零值
	fmt.Println(person{"Kim", 20})

	fmt.Println(person{age: 15, name: "Alice"})

	fmt.Println(person{name: "Alpha"})

	fmt.Println(&person{"Ann", 18})

	s := person{"Sever", 20}
	fmt.Println(s.name)

	sp := new(person)
	fmt.Println(sp.age)

	sp.name = "Pans"
	fmt.Println(sp)
}
