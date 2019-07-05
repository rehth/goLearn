package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "世界，Are You Ok?"
	fmt.Println("Bytes count:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	fmt.Println("Byte Slice:")
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	fmt.Println("String:")
	for i, ch := range s {
		fmt.Printf("(%d %c %X) ", i, ch, ch)
	}
	fmt.Println()

	fmt.Println("Rune Slice:")
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c %X) ", i, ch, ch)
	}
	fmt.Println()

	// strings
	// Fields	Split		Join
	// Contains	Index
	// ToLower	ToUpper
	// Trim		TrimRight	TrimLeft
}
