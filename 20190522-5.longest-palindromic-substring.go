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
		expand := 1
		for mid-expand >= 0 && mid+expand < len(s) && s[mid-expand] == s[mid+expand] {
			expand++
		}
		if expand*2-1 > best {
			best = expand*2 - 1
			bestStr = s[mid-expand+1 : mid+expand]
		}
		// 偶数长度的回文串
		expand = 1
		for mid-expand+1 >= 0 && mid+expand < len(s) && s[mid-expand+1] == s[mid+expand] {
			expand++
		}
		if expand*2-2 > best {
			best = expand*2 - 2
			bestStr = s[mid-expand+2 : mid+expand]
		}
	}
	return bestStr
}
