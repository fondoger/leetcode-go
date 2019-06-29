package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0
	for left < right {
		area := min11(height[left], height[right]) * (right - left)
		best = max11(best, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return best
}

func min11(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max11(a, b int) int {
	if a > b {
		return a
	}
	return b
}
