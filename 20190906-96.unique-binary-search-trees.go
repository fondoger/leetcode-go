package main

/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// 解法2: DP动态规划方式
// 考虑解法1中的递归解，helper(min, max)看作为numTrees(max-min+1)
// dp[2] = dp[0]*dp[1] + dp[1]*dp[0] = 2
// dp[3] = dp[0]*dp[2] + dp[1]*dp[1] + dp[2]*dp[0] = 1*2+1*1+2*1=5
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 解法1: 递归解
func numTrees_Recursively(n int) int {
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
