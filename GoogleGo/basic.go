package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	// 包内部变量
	aa = 3
	bb = true
	ss = "str"
)

func variableZeroValue() {
	// 变量名在变量类型之前
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	// 变量定义时初始化值
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 变量类型推断
	var a, b, c, s = 2, 3, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//	变量定义简写，仅在函数内部使用
	a, b, c, s := 2, 3, true, "def"
	fmt.Println(a, b, c, s)

}

func euler() {
	// 欧拉公式
	fmt.Println(cmplx.Exp(1i*math.Pi)+1, cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	// 类型装换是强制的
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	// 常量的定义，可作为各种类型使用
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, python, golang, javascript)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
}
