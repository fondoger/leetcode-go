package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */
func sum416(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func canPartition(nums []int) bool {
	sum := sum416(nums)
	if sum%2 != 0 {
		return false
	}
	targetSum := sum / 2
	return hasSubArraySumTo(nums, targetSum)
}

// 0-1背包问题, 对于每个数字, 我们有选择或不选择它两种选择
// 是否可以计算的第sum个
// dp[i][j] 表示由[0,i]前缀数组中是否含有子数组的和等于数字j
// 转移方程：dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
func hasSubArraySumTo(nums []int, target int) bool {
	dp := make2Darray(len(nums), target+1)
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}

func make2Darray(m, n int) [][]bool {
	res := make([][]bool, m)
	for i := 0; i < m; i++ {
		res[i] = make([]bool, n)
	}
	return res
}
