package main

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max104(maxDepth(root.Left), maxDepth(root.Right))
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}
