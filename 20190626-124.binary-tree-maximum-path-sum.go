package main

import "math"

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	best := math.MinInt64
	helper124(&best, root)
	return best
}

func helper124(best *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper124(best, root.Left)
	right := helper124(best, root.Right)
	sum := root.Val
	if left > 0 {
		sum += left
	}
	if right > 0 {
		sum += right
	}
	if sum > *best {
		*best = sum
	}
	max := max124(0, max124(root.Val+left, root.Val+right))
	return max
}

func max124(a, b int) int {
	if a > b {
		return a
	}
	return b
}
