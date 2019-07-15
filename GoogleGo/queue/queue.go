package queue

import "errors"

// 一个先进先出队列
type Queue []int

// 将这个元素添加到队列中
// 	q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// 弹出这个队列的头部元素，如果存在的话
func (q *Queue) Pop() (head int, err error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	head = (*q)[0]
	*q = (*q)[1:]
	return
}

// 返回一个 true，如果队列为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
