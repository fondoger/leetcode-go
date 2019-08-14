package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// 方法1: 暴力递归
func wordBreak139_1(s string, wordDict []string) bool {
	hashSet := make(map[string]interface{})
	for _, w := range wordDict {
		hashSet[w] = nil
	}

	var helper139 func(string) bool
	helper139 = func(s string) bool {
		if _, ok := hashSet[s]; ok {
			return true
		}

		for i := 1; i < len(s); i++ {
			_, ok := hashSet[s[:i]]
			if ok && helper139(s[i:len(s)]) {
				return true
			}
		}

		return false
	}

	return helper139(s)
}

// 方法2: 递归 + memo(方法1的改进)
func wordBreak139_2(s string, wordDict []string) bool {
	hashSet := make(map[string]interface{})
	for _, w := range wordDict {
		hashSet[w] = nil
	}
	memo := make(map[string]bool)

	var helper139 func(string) bool
	helper139 = func(s string) bool {
		if val, ok := memo[s]; ok {
			return val
		}
		if _, ok := hashSet[s]; ok {
			memo[s] = true
			return true
		}
		for i := 1; i < len(s); i++ {
			_, ok := hashSet[s[:i]]
			if ok && helper139(s[i:len(s)]) {
				memo[s] = true
				return true
			}
		}
		memo[s] = false
		return false
	}

	return helper139(s)
}

// 方法3: 一维动态规划
// f(i) 表示从[0,i)范围内是否合法
// 则 f(j) = f(i) && Contains(s[i:j))
func wordBreak(s string, words []string) bool {
	hashSet := make(map[string]interface{})
	for _, w := range words {
		hashSet[w] = nil
	}
	contains := func(t string) bool {
		_, ok := hashSet[t]
		return ok
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
