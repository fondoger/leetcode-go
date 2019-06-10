package main

func main() {
	arr := buildIntArray("[-1, 0, 1, 2, -1, -4]")
	for _, item := range threeSum(arr) {
		printIntArray(item)
	}
}
