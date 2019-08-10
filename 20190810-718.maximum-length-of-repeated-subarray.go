package main

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// 提示：
// 1 1 2 3 4
// 1 2 3 4 5
//
// 1. 最长公共子串的写法
func findLength(A []int, B []int) int {
	max := 0
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				if dp[i+1][j+1] > max {
					max = dp[i+1][j+1]
				}
			}
		}
	}
	return max
}

// 2 最长公共子序列的写法
func findLength2(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i+1][j] > dp[i][j+1] {
					dp[i+1][j+1] = dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	return dp[len(A)][len(B)]
}
