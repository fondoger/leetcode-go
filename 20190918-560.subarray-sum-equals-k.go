package main

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// HashMap解法
func subarraySum(nums []int, k int) int {
	var sum, result int
	preSum := make(map[int]int)
	preSum[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if val, ok := preSum[sum-k]; ok {
			result += val
		}
		preSum[sum] += 1
	}

	return result
}

// 暴力解法
func subarraySum_Solution1(nums []int, k int) int {
	dp := make([]int, len(nums)+1)
	var res int
	for i := 1; i <= len(nums); i++ {
		dp[i-1] = 0
		for j := i; j <= len(nums); j++ {
			dp[j] = dp[j-1] + nums[j-1]
			if dp[j] == k {
				res++
			}
		}
	}
	return res
}
