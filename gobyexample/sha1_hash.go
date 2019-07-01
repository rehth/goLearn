package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "Hello World！你好， 世界！"
	// 产生一个散列值得方式是 sha1.New()，同理：可使用 MD5
	h := sha1.New()

	// 写入将 hash 的数据
	h.Write([]byte(s))

	// Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要
	bs := h.Sum(nil)

	fmt.Println("hash前:", s)
	// SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串
	fmt.Printf("hash后: %x\n", bs)
}
