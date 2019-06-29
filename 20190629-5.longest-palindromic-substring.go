/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package main

func longestPalindrome(s string) string {
	best, bestIdx := 0, 0
	for i := 0; i < len(s); i++ {
		// 奇数长度的回文串
		if tmp := spread(&s, i-1, i+1); tmp > best {
			best = tmp
			bestIdx = i - tmp/2
		}
		// 偶数长度的回文串
		if tmp := spread(&s, i, i+1); tmp > best {
			best = tmp
			bestIdx = i - tmp/2 + 1
		}
	}
	return s[bestIdx : bestIdx+best]
}

func spread(s *string, low, high int) int {
	for low >= 0 && high < len(*s) {
		if (*s)[low] != (*s)[high] {
			break
		}
		low--
		high++
	}
	return high - low - 1
}
