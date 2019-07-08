package main

import (
	"fmt"
	"goLearn/GoogleGo/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	// 结构体实例或实例指针都使用 . 来访问属性
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Left = tree.CreateNode(8)

	root.SetValue(6)
	root.Print()

	var pRoot *tree.Node
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.Print()
	pRoot.SetValue(100)
	pRoot.Print()
	root.Print()

	nodes := []tree.Node{
		// 省略 结构体名字
		{Value: 3},
		{},
		{8, nil, &root},
	}
	fmt.Println(nodes)
	pRoot.Traverse()
}
