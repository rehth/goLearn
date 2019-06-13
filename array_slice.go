package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	arr1 := [4]int{}                // 定义一个数组（长度固定、元素类型相同的序列）
	arr2 := [4]int{6, 7, 8}         // 顺序指定数组每个元素
	arr3 := [4]int{3: 4}            // 指定下标的方式
	arr4 := [...]int{1, 3, 5, 7, 9} // 自行推动数组长度
	arr5 := [3][4]int{}             // 支持多维数组创建
	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	slice1 := []int{3, 4, 5}     // 直接创建一个slice（长度可变、元素类型相同的序列）
	slice2 := arr4[:3]           // 通过数组切片创建 slice
	slice3 := make([]int, 4, 10) // 通过构造函数 make 创建，需指定长度，容量
	fmt.Println(slice1, slice2, slice3)

	tick := time.NewTicker(500 * time.Millisecond)
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	go func() {
		for true {
			select {
			case t := <-tick.C:
				fmt.Println("tick:", t)
			case <-cancelCtx.Done():
				fmt.Println("go Done")
				return
			}
		}
		fmt.Println("go end")
	}()
	time.Sleep(time.Second)
	cancelFunc()
	tick.Stop()
	time.Sleep(time.Second)
	fmt.Println("end")
}
