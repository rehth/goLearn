package main

import "fmt"

type rect struct {
	width, height int
}

// 指针类型方法
func (r *rect) area() int {
	return r.width * r.height
}

// 值类型方法
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{12, 15}

	// Go 自动处理方法调用时的值和指针之间的转化
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

}
