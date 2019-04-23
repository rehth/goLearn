package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func main() {
	res := plus(10, 11)
	fmt.Println("10+11=", res)
}
