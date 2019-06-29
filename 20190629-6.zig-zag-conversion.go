package main

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strs := make([]string, numRows)
	row, step := 0, 1
	for i := 0; i < len(s); i++ {
		strs[row] += string(s[i])
		row += step
		if row >= numRows {
			row -= 2
			step = -1
		} else if row < 0 {
			row = 1
			step = 1
		}
	}
	res := ""
	for _, s := range strs {
		res = res + s
	}
	return res
}
