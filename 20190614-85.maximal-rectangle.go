package main

/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		// update heights
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		if largestRectangleArea_85(heights) > res {
			res = largestRectangleArea_85(heights)
		}
	}
	return res
}

// 这个函数与leetcode #84直方图的最大矩形解法一毛一样
func largestRectangleArea_85(heights []int) int {
	n := len(heights)
	lefts := make([]int, n)
	rights := make([]int, n)
	lefts[0] = -1
	rights[n-1] = n
	for i := 1; i < n; i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lefts[p]
		}
		lefts[i] = p
	}
	for j := n - 2; j >= 0; j-- {
		p := j + 1
		for p < n && heights[p] >= heights[j] {
			p = rights[p]
		}
		rights[j] = p
	}
	max := 0
	for i := 0; i < n; i++ {
		cur := (rights[i] - lefts[i] - 1) * heights[i]
		if cur > max {
			max = cur
		}
	}
	return max
}
