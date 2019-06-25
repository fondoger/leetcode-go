package main

func main() {
	if true {
		tree := buildBinaryTree("[3,9,20,null,null,15,7]")
		for _, row := range levelOrderBottom(tree) {
			printIntArray(row)
		}
	}
}
