package main

import "fmt"

func main() {
	tree := buildBinaryTree("[3,5,1,6,2,0,8,null,null,7,4]")
	five := tree.Left
	fore := tree.Left.Right.Right
	fmt.Println(five.Val, fore.Val)
	res := lowestCommonAncestor(tree, five, fore)
	printBinaryTree(res)
}
