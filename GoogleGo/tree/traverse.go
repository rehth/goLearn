package tree

func (node *Node) Traverse() {
	// 中序遍历
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
