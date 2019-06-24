package main

func main() {
	if true {
		t1 := buildBinaryTree("[3,9,20,null,null,15,7]")
		for _, row := range levelOrder(t1) {
			printIntArray(row)
		}
	}
	tmp := []int{1, 2, 3}
	copi := tmp[:]
	println(cap(copi))
}
