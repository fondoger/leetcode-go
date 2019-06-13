package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	fromLeft := make([]int, len(heights))
	fromRight := make([]int, len(heights))
	fromLeft[0] = -1
	fromRight[len(heights)-1] = len(heights)

	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = fromLeft[p]
		}
		fromLeft[i] = p
	}
	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = fromRight[p]
		}
		fromRight[i] = p
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (fromRight[i] - fromLeft[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
