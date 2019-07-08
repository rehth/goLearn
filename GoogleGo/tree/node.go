package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

// 自定义工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 值接收者
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 指针接收者
func (node *Node) SetValue(value int) {
	// 只有指针才可以改变结构体内容
	if node == nil {
		fmt.Println("setting Value to nil node. Ignored")
		return
	}
	node.Value = value
}
