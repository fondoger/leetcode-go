package main

func main() {
	if true {
		root := buildBinaryTree("[1,2,null,3]")
		flatten(root)
		printBinaryTree(root)
	}
}
