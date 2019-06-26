package main

func main() {
	if true {
		root := buildBinaryTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]")
		for _, list := range pathSum(root, 22) {
			printIntArray(list)
		}
	}
}
