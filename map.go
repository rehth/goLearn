package main

import "fmt"


func main () {
	// 字典 map 的创建
	map1 := map[string]string{}				// 使用 map 关键字创建字典（长度不定）
	map2 := map[string]string{"a": "1"}		// 创建时指定元素 k:v
	map3 := make(map[string]string)			// 使用构造函数 make 创建
    map4 := make(map[string]string, 5)		// 创建时指定容量 cap
    fmt.Println(map1, map2, map3, map4)
	
	// 字典 map 的增删改查
	map2["b"] = "2"							// key不存在时新建
	map2["a"] = "a"							// key存在时修改
	delete(map2, "a")					// 删除指定key
	val, flag := map2["b"]					// 根据参数个数来返回数据，如果有两个，则第二个为是否存在的标志
	fmt.Println(map2, val, flag)
	
	// 字典 map 的遍历
	map2["c"] = "3"
	for key, val := range map2 {
		fmt.Println("for:",key, val)
	}
}
