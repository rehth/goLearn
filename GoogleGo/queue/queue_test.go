package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	// Output:
	// 1 <nil>
	// 2 <nil>
	// false
	// 3 <nil>
	// true
	// 0 empty queue

}
