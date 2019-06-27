package main

func main() {
	if true {
		root := buildBinaryTree("[10,5,-3,3,2,null,11,3,-2,null,1]")
		println(pathSum437(root, 8))
	}
	if true {
		root := buildBinaryTree("[1,null,2,null,3,null,4,null,5]")
		println(pathSum437(root, 3))
	}
}
