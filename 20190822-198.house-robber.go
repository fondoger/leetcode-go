package main

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
func rob(nums []int) int {
	pre1, pre2 := 0, 0
	for _, num := range nums {
		if pre1+num > pre2 {
			pre1, pre2 = pre2, pre1+num
		} else {
			pre1, pre2 = pre2, pre2
		}
	}
	return pre2
}
