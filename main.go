package main

func main() {
	if true {
		root := buildBinaryTree("[-10,9,20,null,null,15,7]")
		println(maxPathSum(root))
	}
}
