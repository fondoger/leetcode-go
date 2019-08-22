package main

import "strings"

/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// 超时的暴力解法
func wordBreak_TLE(s string, wordDict []string) []string {
	hash := make(map[string]interface{})
	for _, w := range wordDict {
		hash[w] = nil
	}
	contains := func(s string) bool {
		_, ok := hash[s]
		return ok
	}

	res := make([]string, 0)
	segs := make([]string, 0, len(wordDict)*2)
	helper140(s, contains, &res, segs)
	return res
}

func helper140(s string, contains func(string) bool, res *[]string, segs []string) {
	if s == "" {
		if len(segs) > 0 {
			*res = append(*res, strings.Join(segs, " "))
		}
		return
	}
	t := len(segs)
	segs = append(segs, "")
	for i := 1; i <= len(s); i++ {
		if contains(s[:i]) {
			segs[t] = s[:i]
			helper140(s[i:], contains, res, segs)
		}
	}
}

// 假设wordDict很小，我们可以可以每次遍历wordDict列表，而非拆分原始字符串
// 这样的话遇到极端case的情况比较小
func wordBreak2(s string, wordDict []string) []string {
	return dfs140(s, wordDict, make(map[string][]string))
}

// m 存放每个字符串可以划分的小串
func dfs140(s string, words []string, m map[string][]string) []string {
	if val, ok := m[s]; ok {
		return val
	}

	res := make([]string, 0)
	if len(s) == 0 { // 挺重要的哈
		res = append(res, "")
		m[s] = res
		return res
	}

	for _, w := range words {
		if strings.HasPrefix(s, w) {
			sublist := dfs140(s[len(w):], words, m)
			for _, substr := range sublist {
				res = append(res, w+AorB(len(substr) == 0, "", " ")+substr)
			}
		}
	}

	m[s] = res
	return res
}

func AorB(val bool, a string, b string) string {
	if val {
		return a
	}
	return b
}
