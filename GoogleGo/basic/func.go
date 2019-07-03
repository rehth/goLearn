package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation：%s", op)
	}
}

// 多值返回
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 使用函数充当函数参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function <%s> with args (%d %d)\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) (res int) {
	for _, v := range numbers {
		res += v
	}
	return res
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}

// 更合理的方式
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	res, err := eval(3, 4, "x")
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
	q, r := div(15, 4)
	fmt.Println(q, r)

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	//swap1(&a, &b)
	a, b = swap2(a, b)
	fmt.Println(a, b)
}
