/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package main

/* Solution 1: */
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstCharMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return firstCharMatch && isMatch(s[1:], p) || isMatch(s, p[2:])
	} else {
		return firstCharMatch && isMatch(s[1:], p[1:])
	}
}
