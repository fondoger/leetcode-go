/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	count := 0
	error := false
	for count < len(strs[0]) {
		for s := 1; s < len(strs); s++ {
			if count >= len(strs[s]) || strs[s][count] != strs[0][count] {
				error = true
				break
			}
		}
		if !error {
			count++
		} else {
			break
		}
	}
	return strs[0][:count]
}
