package main

func main() {
	if true {
		arr := buildIntArray("[2,1,2]")
		for _, arr := range permuteUnique(arr) {
			printIntArray(arr)
		}
	}
}
