package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err.Error())
		} else {
			panic(fmt.Sprintln("I don't know what to do:", r))
		}
	}()
	panic(123)
	//a, b := 5, 0
	//fmt.Println(a / b)
	//panic(errors.New("this is an error"))
}
func main() {
	tryRecover()
}
