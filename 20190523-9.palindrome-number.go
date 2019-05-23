/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */
package main

/* Solution 1 */
// 注意：这题用不到位运算
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	originX := x
	reverseX := 0
	digits := uint(0)
	for x > 0 {
		reverseX = reverseX*10 + x%10
		x = x / 10
		digits++
	}
	return originX == reverseX
}
