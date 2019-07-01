package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//  os.Setenv 来设置一个键值队。使用 os.Getenv获取一个键对应的值
	os.Setenv("Foo", "bar")
	fmt.Println("Foo", os.Getenv("Foo"))
	fmt.Println()
	// os.Environ 来列出所有环境变量键值队
	for _, value := range os.Environ() {
		fmt.Println(strings.Split(value, "="))
	}
}
