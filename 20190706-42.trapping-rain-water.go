package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
func trap(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		if height[left] < height[right] {
			top := height[left]
			left++
			for left < right && height[left] <= top {
				res += top - height[left]
				left++
			}
		} else {
			top := height[right]
			right--
			for left < right && height[right] <= top {
				res += top - height[right]
				right--
			}
		}
	}
	return res
}
