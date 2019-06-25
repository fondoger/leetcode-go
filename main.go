package main

func main() {
	if true {
		pre := buildIntArray("[3,9,20,15,7]")
		in := buildIntArray(" [9,3,15,20,7]")
		tree := buildTree(pre, in)
		printBinaryTree(tree)
	}
}
