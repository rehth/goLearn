package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机数种子是确定的，使每次产生相同的随机数序列
	// 伪随机数：Int  0 <= n <= 100
	fmt.Println(rand.Intn(100), rand.Intn(100))
	// 随机浮点数：float64  0.0 <= f <= 1.0
	fmt.Println((rand.Float64()*5)+5, (rand.Float64()*5)+5)

	// 给定一个变化的种子，产生 source 对象
	source := rand.NewSource(time.Now().UnixNano())
	// New(): 返回一个新的随机数生成器（伪） *rand
	random := rand.New(source)
	// 新的 *rand 同 rand 用法一致
	fmt.Println(random.Intn(100), random.Intn(100))
	fmt.Println((rand.Float64()*5)+5, (rand.Float64()*5)+5)

}
