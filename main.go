package main

func main() {
	if true {
		pre := buildIntArray("[9,3,15,20,7]")
		in := buildIntArray(" [9,15,7,20,3]")
		tree := buildTree(pre, in)
		printBinaryTree(tree)
	}
}
