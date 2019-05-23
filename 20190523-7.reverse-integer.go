/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
package main

/* Solution 1: O(n) */
func reverse(x int) int {
	n := int32(x) // golang中的int可能是64位的
	res := int32(0)
	for n != 0 {
		if (res*10+n%10)/10 != res {
			return 0 // overflow
		}
		res = res*10 + n%10
		n /= 10
	}
	return int(res)
}
