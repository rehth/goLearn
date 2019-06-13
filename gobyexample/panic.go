package main

import (
	"fmt"
	"os"
)

func main() {
	panic("One problem arises here")

	fmt.Println("我想继续的操作")

	f, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
