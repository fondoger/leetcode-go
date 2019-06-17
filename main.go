package main

func main() {
	arr := buildIntArray("[1,2,3]")
	for _, arr := range(permute(arr)) {
		printIntArray(arr)
	}
}
