package container

import "fmt"

func updateSlice(s []int, ind int) {
	s[ind] = 100
}

func printSplit() {
	fmt.Println("***************************************")
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[2:]  =", arr[2:])
	fmt.Println("arr[:6]  =", arr[:6])
	fmt.Println("arr[:]   =", arr[:])
	printSplit()
	s1, s2 := arr[2:6], arr[2:]
	updateSlice(s1, 0)
	fmt.Println("update slice1 for 0:", s1, s2)
	updateSlice(s2, 3)
	fmt.Println("update slice2 for 3:", s1, s2)
	fmt.Println("arr", arr)
	printSplit()
	// slice 可以向后扩展，但不可向前扩展
	// s[i]不可超越len(s)，扩展不可超越底层数组cap(s)
	fmt.Println("Extending slice")
	arr[2], arr[5] = 2, 5
	fmt.Println("arr", arr)
	s1 = arr[2:6]
	s3 := s1[3:5]
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s3=%v len(s3)=%d cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Println("s3[:4] =", s3[:4])
	printSplit()
	// append 元素超越cap时，系统会重新分配底层数组
	s4 := append(s3, 10)
	s5 := append(s4, 11)
	s6 := append(s5, 12)
	fmt.Println("s4, s5, s6 =", s4, s5, s6)
	fmt.Println("arr=", arr)
	printSplit()
	var s []int // zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println("s =", s)
	printSplit()
	fmt.Println("Creating slice")
	s7 := make([]int, 10, 34)
	fmt.Printf("s7=%v len(s7)=%d cap(s7)=%d\n", s7, len(s7), cap(s7))

	fmt.Println("Copying slice")
	copy(s7, s5)
	fmt.Printf("s7=%v len(s7)=%d cap(s7)=%d\n", s7, len(s7), cap(s7))

	fmt.Println("Delete element from slice")
	s8 := append(s7[:3], s7[4:]...)
	fmt.Printf("s8=%v len(s8)=%d cap(s8)=%d\n", s8, len(s8), cap(s8))

}
