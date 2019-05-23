/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */
package main

/* Solution 1: using map */
func intToRoman(num int) string {
	res := ""
	mapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, key := range keys {
		for num >= key {
			res += mapping[key]
			num -= key
		}
	}
	return res
}
