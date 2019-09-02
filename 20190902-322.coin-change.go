package main

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// 动态规划, 以数量从1到amount依次算出来
func coinChange_SOLUTION_2(coins []int, amount int) int {
	// dp[amount] 记录转换次数
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = -1
	}
	for m := 1; m <= amount; m++ {
		for _, c := range coins {
			if m-c >= 0 {
				dp[m] = best322(dp[m-c]+1, dp[m])
			}
		}
	}
	return dp[amount]
}

func best322(a, b int) int {
	if a == 0 {
		return b
	}
	if b == -1 {
		return a
	} else if a <= b {
		return a
	}
	return b
}

// 解法1: 自顶向下的memo数组+暴力搜索
func coinChange_SOLUTION_1(coins []int, amount int) int {
	m := make(map[int]int, 0)
	return helper322(m, coins, amount)
}

func helper322(m map[int]int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}
	if val, ok := m[amount]; ok {
		return val
	}

	res := -1
	for _, c := range coins {
		val := helper322(m, coins, amount-c)
		if val != -1 {
			if res == -1 || val < res {
				res = val
			}
		}
	}

	if res != -1 {
		m[amount] = res + 1
	} else {
		m[amount] = -1
	}
	return m[amount]
}
