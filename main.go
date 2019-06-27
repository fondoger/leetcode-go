package main

func main() {
	if true {
		matrix := [][]int{
			buildIntArray("[1,2,3]"),
			buildIntArray("[4,5,6]"),
			buildIntArray("[7,8,9]"),
		}
		rotate(matrix)
	}
}
