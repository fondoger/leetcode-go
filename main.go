package main

func main() {
	if true {
		t1 := buildBinaryTree("[3,9,20,null,null,15,7]")
		for _, row := range zigzagLevelOrder(t1) {
			printIntArray(row)
		}
	}
}
