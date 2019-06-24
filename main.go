package main

func main() {
	if true {
		t1 := buildBinaryTree("[1,2,3]")
		t2 := buildBinaryTree("[1,2,3]")
		println(isSameTree(t1, t2))
	}
	if true {
		t1 := buildBinaryTree(" [1,2]")
		t2 := buildBinaryTree(" [1,null,2]")
		println(isSameTree(t1, t2))
	}
	if true {
		t1 := buildBinaryTree("    [1,2,1]")
		t2 := buildBinaryTree("[1,1,2]")
		println(isSameTree(t1, t2))
	}
}
