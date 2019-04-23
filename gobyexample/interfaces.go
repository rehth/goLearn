package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// 结构体类型 rectangle 实现了 geometry接口
type rectangle struct {
	width, height float64
}

// 结构体类型 circle 实现了 geometry接口
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g, g.area(), g.perim())
}

func main() {
	rc := rectangle{height: 5, width: 4}
	cc := circle{3}
	// 实现某个接口后可做该类型
	measure(rc)
	measure(cc)
}
