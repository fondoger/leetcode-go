package main

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
func numTrees(n int) int {
	return helper96(1, n)
}

func helper96(low, high int) int {
	if low >= high {
		return 1
	}
	res := 0
	for root := low; root <= high; root++ {
		left := helper96(low, root-1)
		right := helper96(root+1, high)
		res += left * right
	}
	return res
}
