package main

func main() {
	if true {
		t1 := buildBinaryTree("[1,2,2,3,4,4,3] ")
		println(isSymmetric(t1))
	}
	if true {
		t1 := buildBinaryTree("[1,2,2,null,3,null,3] ")
		println(isSymmetric(t1))
	}

}
