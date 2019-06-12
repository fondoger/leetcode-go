package main

/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

func min_424(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	max := 0
	for i := 0; i < len(s); i++ {
		if i != 0 && s[i] == s[i-1] {
			continue
		}
		j := i + 1
		remain := k
		for remain >= 0 && j < len(s) {
			if s[j] != s[i] {
				if remain == 0 {
					break
				}
				remain--
			}
			j++
		}
		size := j - i
		if j == len(s) {
			size += min_424(i, remain)
		}
		if size > max {
			max = size
		}
	}
	return max
}
