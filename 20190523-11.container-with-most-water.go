/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	best := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > best {
			best = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return best
}
