package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	p, _ := os.Getwd()
	fname := p + "\\const.go"
	// ioutil 包一次读取文件所有数据
	fmt.Println(p, fname)
	dat, err := ioutil.ReadFile(fname)
	check(err)
	fmt.Printf("ioutil.ReadFile：\n%s\n", string(dat))

	// os 读取文件
	f, err := os.Open(fname)
	defer f.Close()
	check(err)
	bs := make([]byte, 15)
	ns, err := f.Read(bs)
	check(err)
	fmt.Printf("%d bytes: %s\n", ns, string(bs))

	// io 包提供了一些可以帮助我们进行文件读取的函数
	sk, err := f.Seek(15, 0)
	check(err)
	ss := make([]byte, 5)
	ks, err := io.ReadAtLeast(f, ss, 1)
	check(err)
	fmt.Printf("%d bytes @ %d:%s\n", ks, sk, string(ss))

	// seek 移动文件指针
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 包实现了带缓冲的读取
	br := bufio.NewReader(f)
	bb, err := br.Peek(20)

	check(err)
	fmt.Printf("%d bytes: %s\n", len(bb), string(bb))
	f.Close()
}
