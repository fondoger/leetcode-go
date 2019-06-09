package main

func main() {
	t := "[5,4,7,3,null,2,null,-1,null,9]"

	tree := buildBinaryTree(t)
	printBinaryTree(tree)

	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Left.Left = &TreeNode{-1, nil, nil}
	root.Right = &TreeNode{7, nil, nil}
	root.Right.Left = &TreeNode{2, nil, nil}
	root.Right.Left.Left = &TreeNode{9, nil, nil}
	printBinaryTree(root)
}
