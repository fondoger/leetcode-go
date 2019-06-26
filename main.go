package main

func main() {
	if true {
		root := buildBinaryTree("[3,5,1,6,2,0,8,null,null,7,4]")
		five := root.Left
		four := root.Left.Right.Right
		println(lowestCommonAncestor(root, five, four).Val)
	}
}
