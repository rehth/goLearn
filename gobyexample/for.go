package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // 类似 Java for 循环语句
		fmt.Println("for:", i)
	}

	i := 10
	for i > 0 { // 类似于 Python while 语句
		fmt.Println("while:", i)
		i = i - 2
	}

	for { // 类似于 Python while True
		if i > 10 {
			break
		}
		fmt.Println("while True:", i)
		i++
	}

	arr := [...]int{6, 7, 8, 9}
	for ind, el := range arr { // for-range：遍历可迭代对象
		fmt.Println("enumerate:", ind, el) // 此时 range 相当于 Python enumerate 关键字
	}
}
