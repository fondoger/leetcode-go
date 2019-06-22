package main

func main() {
	if true {
		arr1 := buildIntArray("[1,2,3,4,5]")
		arr2 := buildIntArray("[4,5,3,2,1]")
		println(validateStackSequences(arr1, arr2))
	}
	if true {
		arr1 := buildIntArray("[1,2,3,4,5]")
		arr2 := buildIntArray("[4,3,5,1,2]")
		println(validateStackSequences(arr1, arr2))
	}
}
