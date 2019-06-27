package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

// 结构字段声明标签来自定义编码的 JSON 数据键名称
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	fmt.Println("json dump：.................................")
	// 基本类型转 JSON
	bolB, _ := json.Marshal(true)
	fmt.Println("bool:", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("int:", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("float:", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("string:", string(strB))

	// go 容器对象转 JSON
	slcB, _ := json.Marshal([3]string{"apple", "peach", "pear"})
	fmt.Println("slice/arr:", string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 5, "peach": 3, "pear": 4})
	fmt.Println("map:", string(mapB))

	// 自定义类型转 JSON
	res1, _ := json.Marshal(Response1{1, []string{"apple", "pear"}})
	fmt.Println("response1:", string(res1))

	res2, _ := json.Marshal(Response2{1, []string{"peach", "pear"}})
	fmt.Println("response2:", string(res2))

	fmt.Println("json load：.................................")
	byt := []byte(`{"num":6.13,"slice":["a","b"]}`)
	var dat map[string]interface{}
	// 实际的解码和相关的错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("bytes：%c %v\n", byt, dat)

	// 进行适当的类型转换，如：将 num 的值转换成 float64类型
	num, ok := dat["num"].(float64)
	fmt.Println(num, ok)

	// 嵌套的值需要一系列的转化
	slice := dat["slice"].([]interface{})
	str1 := slice[0].(string)
	fmt.Println(str1)

	// 可解码 JSON 值到自定义类型
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 将 JSON 编码直接输出至 os.Writer流中
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
