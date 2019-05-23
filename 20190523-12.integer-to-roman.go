/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
package main

/* Solution 1: using map */
func intToRoman(num int) string {
	res := ""
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	values := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i, key := range keys {
		for num >= key {
			res += values[i]
			num -= key
		}
	}
	return res
}
