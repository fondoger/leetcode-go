package main

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

func max213(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob_REMOVE_THIS(nums []int) int {
	a := robFrom(nums, 0)
	b := robFrom(nums, 1)
	c := robFrom(nums, 2)
	return max213(a, max213(b, c))
}

func robFrom(nums []int, start int) int {
	if start >= len(nums) { // 注意点1: 边界判断
		return 0
	}
	l, r := 0, nums[start] // 注意点2: 起始点必须包含
	for i := start + 1; i < start+len(nums)-1; i++ {
		if nums[i%len(nums)]+l > r {
			l, r = r, nums[i%len(nums)]+l
		} else {
			l, r = r, r
		}
	}
	return r
}
