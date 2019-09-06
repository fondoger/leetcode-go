package main

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
func numTrees(n int) int {
	return helper(1, n)
}

func helper(min, max int) int {
	if min >= max {
		return 1
	}
	res := 0
	for i := min; i <= max; i++ {
		res += helper(min, i-1) * helper(i+1, max)
	}
	return res
}
