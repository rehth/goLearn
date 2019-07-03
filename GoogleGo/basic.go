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
	// 变量定义简写，仅在函数内部使用
	a, b, c, s := 2, 3, true, "def"
	fmt.Println(a, b, c, s)

}

func euler() {
	// 欧拉公式
	fmt.Println(cmplx.Exp(1i*math.Pi)+1, cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	// 类型转换是强制的
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
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
}
