package main

import "fmt"

type Stu struct {
	name 	string
	age  	uint8
	score 	[3]int
}

func main () {
	stu1 := Stu{"zhang", 16, [3]int{98, 95, 72}}				// 顺序参数
	stu2 := Stu{name:"li", score:[3]int{87, 96, 84}, age:17}	// 关键字参数
	stu3 := new(Stu)											// 使用 new 来给该结构体变量分配内存，
	stu3.name = "wang"											// 它返回指向已分配内存的指针。
	stu3.age = 18
	fmt.Println(stu1, stu2, stu3.score)
}
																
