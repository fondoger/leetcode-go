package main

func main() {
	if true {
		root := &TreeNode{1, nil, nil}
		root.Left = &TreeNode{2, nil, nil}
		root.Left.Right = &TreeNode{4, nil, nil}
		root.Right = &TreeNode{3, nil, nil}
		root.Right.Left = &TreeNode{1, nil, nil}
		root.Right.Right = &TreeNode{5, nil, nil}
		printBinaryTree(root)
	}
}
