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

		for i := 1; i < len(s)-1; i++ {
			_, ok := hashSet[s[:i]]
			if ok && helper139(s[i:len(s)]) {
				return true
			}
		}

		return false
	}

	return helper139(s)
}

// 方法2: 递归 + memo

func wordBreak(s string, wordDict []string) bool {
	hashSet := make(map[string]interface{})
	for _, w := range wordDict {
		hashSet[w] = nil
	}
	memo := make(map[string]bool)
}
