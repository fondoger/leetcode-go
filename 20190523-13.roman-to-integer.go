/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */
package main

/* Solution 1: */
func romanToInt(s string) int {
	res := 0
	mapping := [128]int{}
	mapping['I'] = 1
	mapping['V'] = 5
	mapping['X'] = 10
	mapping['L'] = 50
	mapping['C'] = 100
	mapping['D'] = 500
	mapping['M'] = 1000
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && mapping[s[i]] < mapping[s[i+1]] {
			res += mapping[s[i+1]] - mapping[s[i]]
			i++
		} else {
			res += mapping[s[i]]
		}
	}
	return res
}
