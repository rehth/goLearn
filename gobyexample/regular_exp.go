package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	re, _ := regexp.Compile("p([a-z]+)ch")

	// 匹配字符串：-> bool
	fmt.Println(re.MatchString("peach"))

	// 查找字符串（索引）：-> str(slice(start, end))
	fmt.Println("FindString:", re.FindString("peach punch"))
	fmt.Println("FindStringIndex:", re.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配（索引位置）
	fmt.Println("FindStringSubMatch:", re.FindStringSubmatch("peach punch"))
	fmt.Println("FindStringSubMatchIndex:", re.FindStringSubmatchIndex("peach punch"))

	// 返回所有的匹配项，可通过 n 限制匹配次数
	fmt.Println("FindAllSting:", re.FindAllString("peach punch pinch", -1))
	fmt.Println("FindAllStringSubMatchIndex:", re.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println("FindAllStingN:", re.FindAllString("peach punch pinch", 2))

	// 可以提供 []byte参数并将 String 从函数命中去掉
	fmt.Println("Match:", re.Match([]byte("peach")))

	// 创建正则表示式常量时，可以使用 Compile 的变体MustCompile
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("Regexp struct:", r)
	// regexp 包也可以用来替换部分字符串为其他值
	fmt.Println("Replace String:", r.ReplaceAllString("a peach punch", "foo"))

	fmt.Println("Replace Func:", string(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)))

}
