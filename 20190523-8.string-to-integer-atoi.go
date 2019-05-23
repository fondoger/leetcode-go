/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
package main

import (
	"math"
	"unicode"
)

/* Solution 1: */
/*
要点：
1. 空白字符只能出现在开头
*/
func myAtoi(str string) int {
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	sign := int64(1)
	if len(str) > 0 && (str[0] == '+' || str[0] == '-') {
		if str[0] == '-' {
			sign = -1
		}
		str = str[1:]
	}
	res := int64(0)
	for _, char := range str {
		if !unicode.IsDigit(char) {
			break
		}
		res = res*10 + (int64(char) - '0')
		if sign == 1 && res > math.MaxInt32 {
			return int(math.MaxInt32)
		} else if sign == -1 && res > math.MaxInt32+1 {
			return int(math.MinInt32)
		}
	}
	return int(res * sign)
}
