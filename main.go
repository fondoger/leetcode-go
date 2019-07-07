package main

func main() {
	if true {
		nums := [][]int{
			{1, 3}, {2, 6}, {8, 10}, {15, 18},
		}
		for _, val := range merge(nums) {
			printIntArray(val)
		}
	}
}
