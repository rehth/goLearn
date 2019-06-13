package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	format := fmt.Printf
	p := point{1, 2}
	format("%v\n", p)

	// %+v 的格式化输出内容将包括结构体的字段名
	format("%+v\n", p)

	// %#v 形式则输出这个值的 Go 语法表示
	format("%#v\n", p)

	// 格式化类型，使用 %T
	format("%T\n", p)

	// 格式化布尔值 %t
	format("%t\n", true)

	//  %d进行标准的十进制格式化
	format("%d\n", 1234)

	// %b %x 提供二进制、十六进制编码， %c 对应ASCII码
	format("%b\n", 12)
	format("%x\n", 12)
	format("%c\n", 66)

	// %f 进行最基本的十进制格式化
	format("%f\n", 12.12)

	// %e 和 %E 将浮点型格式化为 科学技科学记数法
	format("%e\n", 0.3)
	format("%E\n", 12340000.0)

	//%s 进行基本的字符串输出  带有双引号的输出，使用 %q
	format("%s\n", "\"string\"")
	format("%q\n", "\"string\"")

	// %x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示
	format("%x\n", "hex")

	// 输出一个指针的值，使用 %p
	format("%p\n", &p)

	// % 后面使用数字来控制输出宽度 默认右对齐、空格填充
	format("|%6d|%6d|\n", 12, 345)

	// 指定浮点型 宽度.精度 的语法来指定输出
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// 左对齐，使用 - 标志
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Sprintf 则格式化并返回一个字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf 来格式化并输出到指定的 io.Writers
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
