package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Ending...")
	// 使用 os.Exit 时 defer 将不会 执行
	os.Exit(3)
}
