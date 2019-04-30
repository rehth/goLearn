package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if num == 88 {
		// errors.New 构造一个使用给定的错误信息的基本error 值
		return -1, errors.New("error for f1")
	}
	return num + 3, nil
}

func f2(num int) (int, error) {
	if num == 88 {
		return -1, &numError{num, "error for f2"}
	}
	return num + 3, nil
}

type numError struct {
	num  int
	prob string
}

// 通过实现 Error 方法来自定义 error 类型
func (e numError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

func main() {
	for _, i := range []int{7, 42, 88} {
		// 在 if行内的错误检查代码, 返回错误值为 nil 代表没有错误
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42, 88} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// 通过类型断言来得到这个错误类型的实例
	_, e := f2(88)
	if ne, ok := e.(*numError); ok {
		fmt.Println(ne.num)
		fmt.Println(ne.prob)
	}
}
