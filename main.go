package main

func main() {
	if true {
		root := buildBinaryTree("[1,3,null,null,2]")
		printBinaryTree(root)
		recoverTree(root)
		printBinaryTree(root)
	}
}
