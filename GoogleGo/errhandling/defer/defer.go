package main

import (
	"bufio"
	"fmt"
	"os"
)

type slice []int

func (s *slice) add(element int) *slice {
	*s = append(*s, element)
	fmt.Println(element)
	return s
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// defer 在函数结束时发生
func tryDefer() {
	for i := 0; i < 100; i++ {
		// 参数在 defer 语句时计算
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}

}

// defer: 先进后出（FILO）
func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 666)
	if err != nil {
		// 处理已知错误
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Printf("ERROR: %s %s %s\n", pathErr.Op, pathErr.Path, pathErr.Err)
			// 处理未知错误
		} else {
			panic(err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fib := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fib())
	}
}

func main() {
	//tryDefer()
	writeFile("defer.txt")

	s := make(slice, 5)
	defer s.add(1).add(2)
	s.add(3)

}
