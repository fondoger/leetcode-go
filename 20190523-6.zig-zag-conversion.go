/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strs := make([]string, numRows)
	row, step := 0, 1
	for _, c := range s {
		strs[row] += string(c)
		row += step
		if row == numRows {
			row = numRows - 2
			step = -1
		} else if row == -1 {
			row = 1
			step = 1
		}
	}
	res := ""
	for _, str := range strs {
		res += str
	}
	return res
}
