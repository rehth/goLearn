package tree

func (node *Node) Traverse() {
	// 中序遍历
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	// 中序遍历
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
