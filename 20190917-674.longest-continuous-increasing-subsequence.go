package main

/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := nums[0]
	count, maxCount := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > prev {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 1
		}
		prev = nums[i]
	}
	return maxCount
}
