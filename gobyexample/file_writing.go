package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkerr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	writeString := "测试n "
	filename := "output.txt"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	checkerr(err)
	defer f.Close()

	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	n, err1 := io.WriteString(f, writeString) //写入文件(字符串)
	checkerr(err1)
	fmt.Printf("写入 %d 个字节\n", n)

	/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	d1 := []byte(writeString)
	err2 := ioutil.WriteFile(filename, d1, 0666) //写入文件(字节数组)
	checkerr(err2)

	/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	n2, err3 := f.Write(d1) //写入文件(字节数组)
	checkerr(err3)
	fmt.Printf("写入 %d 个字节\n", n2)
	n3, err3 := f.WriteString("writer\n") //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节\n", n3)
	f.Sync()

	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString("buffer edn")
	fmt.Printf("写入 %d 个字节\n", n4)
	w.Flush()
	f.Close()
}
