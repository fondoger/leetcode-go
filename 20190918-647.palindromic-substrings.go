package main

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += spread(s, i, i+1) / 2
		res += (spread(s, i-1, i+1) + 1) / 2
	}
	return res
}

//
func spread(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}
