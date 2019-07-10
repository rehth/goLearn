package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 闭包实现 累加
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 函数式编程实现 累加
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func fibonacci() fibGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibGen func() int

func (fib fibGen) Read(p []byte) (n int, err error) {
	next := fib()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func printFile(f io.Reader) {
	scanner := bufio.NewScanner(f)
	// while exp:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	add := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, add = add(i)
		fmt.Printf("add %d = %d\n", i, s)
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
	printFile(fib)
}
