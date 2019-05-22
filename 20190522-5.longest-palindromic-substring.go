/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package main

func longestPalindrome(s string) string {
	best := 0
	bestStr := ""
	for mid := 0; mid < len(s); mid++ {
		// 奇数长度的回文串
		spread(&s, len(s), mid-1, mid+1, &best, &bestStr)
		// 偶数长度的回文串
		spread(&s, len(s), mid, mid+1, &best, &bestStr)
	}
	return bestStr
}

func spread(s *string, size, begin, end int, best *int, bestStr *string) {
	for begin >= 0 && end < size && (*s)[begin] == (*s)[end] {
		begin--
		end++
	}
	if end-begin-1 > *best {
		(*best) = end - begin - 1
		(*bestStr) = (*s)[begin+1 : end]
	}
}
