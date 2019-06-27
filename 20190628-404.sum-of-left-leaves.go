package main

/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	helper404(&sum, root, false)
	return sum
}
func helper404(sum *int, root *TreeNode, left bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if left {
			*sum += root.Val
		}
		return
	}
	helper404(sum, root.Left, true)
	helper404(sum, root.Right, false)
}
