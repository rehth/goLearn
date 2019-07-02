package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	// while exp:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// while True:
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		// 0 需要特殊处理
		convertToBin(0),
	)
	printFile("output.txt")
	forever()
}
