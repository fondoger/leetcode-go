package main

func main() {
	if true {
		nums := buildIntArray("[-1, 0, 1, 2, -1, -4]")
		for _, row := range threeSum(nums) {
			printIntArray(row)
		}
	}
}
