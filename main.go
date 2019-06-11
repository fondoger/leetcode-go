package main

func main() {
	arr := buildIntArray("[1,3,-1,-3,5,3,6,7]")
	printIntArray(arr)
	res := maxSlidingWindow(arr, 3)
	printIntArray(res)
}
