package main

import (
	"math"
)

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	minLength, minStartPos := math.MaxInt32, -1
	needs := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	remain := len(t)
	low, high := 0, 0
	for high < len(s) {
		if _, ok := needs[s[high]]; ok { // 确保字符是需要的字符
			if needs[s[high]] > 0 {
				remain--
			}
			needs[s[high]]-- // 如果字符多次出现会变成负数
			for remain == 0 {
				// low开始往前移动
				if _, ok2 := needs[s[low]]; ok2 { // 如果字符是需要的字符
					if high-low < minLength { // 修正最优值
						minLength = high - low + 1
						minStartPos = low
					}
					needs[s[low]]++
					if needs[s[low]] > 0 {
						remain++
					}
				}
				low++
			}
		}
		high++
	}
	if minStartPos == -1 {
		return ""
	}
	return s[minStartPos : minStartPos+minLength]
}
