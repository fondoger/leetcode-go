package main

func main() {
	if true {
		root := buildBinaryTree("[1,null,2,3]")
		printIntArray(inorderTraversal(root))
	}
}
