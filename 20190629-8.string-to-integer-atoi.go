package main

import (
	"math"
	"unicode"
)

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// 审题：空白字符只能出现在开头
// 小技巧：利用golang的string切片，可以很方便的处理

func myAtoi(str string) int {
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	sign := int64(1)
	if len(str) > 0 && (str[0] == '-' || str[0] == '+') {
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
			return math.MaxInt32
		} else if sign == -1 && res > math.MaxInt32+1 {
			return math.MinInt32
		}
	}
	return int(sign * res)
}
