package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Go 提供了一个 flag 包，支持基本的命令行标志解析
	fmt.Println("args:", os.Args)

	strOpt := flag.String("str", "hello", "str 参数的帮助信息。。。")
	boolOpt := flag.Bool("bool", true, "bool 参数的帮助信息。。。")

	var intVar int
	flag.IntVar(&intVar, "int", 10, "int 参数的帮助信息")

	flag.Parse()

	fmt.Println("str:", *strOpt)
	fmt.Println("bool:", *boolOpt)
	fmt.Println("int:", intVar)
	fmt.Println("tail:", flag.Args())

}
