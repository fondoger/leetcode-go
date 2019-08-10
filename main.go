package main

func main() {
	if true {
		A := buildIntArray("[1,2,3,2,1]")
		B := buildIntArray("[3,2,1,4,7]")
		println(findLength(A, B))
	}
	if true {
		A := buildIntArray("[0,0,0,0,1]")
		B := buildIntArray("[1,0,0,0,0]")
		println(findLength(A, B) == 4)
	}
}
