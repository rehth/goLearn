package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

// 自定义工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

// 值接收者
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 指针接收者
func (node *treeNode) setValue(value int) {
	// 只有指针才可以改变结构体内容
	if node == nil {
		fmt.Println("setting value to nil node. Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	// 中序遍历
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	// 结构体实例或实例指针都使用 . 来访问属性
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.left = createNode(8)

	root.setValue(6)
	root.print()

	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.print()
	pRoot.setValue(100)
	pRoot.print()
	root.print()

	nodes := []treeNode{
		// 省略 结构体名字
		{value: 3},
		{},
		{8, nil, &root},
	}
	fmt.Println(nodes)
	pRoot.traverse()
}
