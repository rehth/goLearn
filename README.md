# goLearn
golang learning...
#### Go 简介
> Go 编程语言是开源项目。
> 
> GO的并行机制使其很容易编写多核和网络应用，而新奇的类型系统允许构建有弹性的模块化程序。
> 
> Go 编译到机器码非常快速，同时具有便利的垃圾回收和强大的运行时反射。
> 
> GO是快速的、静态类型编译语言，但是感觉上是动态类型的，解释型语言。

#### GO 的环境搭建及配置
详细情况参考： [CSDN-[Linux]Windows下搭建go语言开发环境及开发IDE](https://note.youdao.com/)

#### Hello World

```go
// 测试我的第一个 golang 程序

/* Hello World!你好世界！ */

// 所有的 Go 文件以 package <something>
// 独立运行的执行文件必须是 package main
// 1. 当前程序的包名
package main

// 2. 导入其他包
import io "fmt"

// 3. 常量的定义
const Pi = 3.14

// 4. 全局变量的定义
var name = "king"

// 5. 一般类型的声明
type ageType int

// 6. 结构的声明
type tree struct{}

// 7. 接口的声明
type golang interface{}

// 8. 使用 main() 作为程序的入口
func main() {
	// 字符串由 " 包裹，并且可以包含非 ASCII 的字符
	io.Println("hello world!你好，世界")
}
```
**注意**：当标识符(包括常量、变量、类型、函数名、结构字段等)以一个大写字母开头，如：Group1，那么该对象可被外部引用，否则，则在外部不可见。(**可见性规则**)

#### Go 关键字及标识符
- **注释**
    > golang 支持单行注释及多行注释
    
    ```go
    // 单行注释
    
    /*
    多行注释
    */
    ```

- **标识符**
    > 由一个或多个字母(`a-zA-Z`)、数字(`0-9`)、下划线(`_`)组成的序列。但不能与关键字相同且第一个字符不能是数字
    
    ```
    无效标识符示例：
    1ab     (以数字开头)
    case    (Go 语言关键字)
    a+b     (不被允许的特殊符号)
    ```

- **关键字**
    - Go 代码中会使用到25个关键字或保留字：

    break | default | func | interface | select
    ---|---|---|---|---
    case | defer | go | map | struct
    chan | else | goto | package | switch
    const | fallthrough | if | range | type
    continue | for | import | return | var
    - GO 还有36个预定义标识符，其中包括了基本类型的名称和一些内置函数

    append | bool | byte | cap | colse | complex
    ---|---|---|---|---|---
    complex64 | complex128 | copy | false | float32 | float64
    imag | int | int8 | int16 | int32 | int64
    iota | len | make | new | nil | panic 
    print | println | real | recover | string | true
    uint | uint8 | uint16 | uint32 | uint64 | uintptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。

除此之外，程序可能使用到这些分隔符：括号`()`、中括号`[]`、大括号`{}`，也可能会使用这些标点符号：`.`、`,`、`;`、`:`和`...`。
