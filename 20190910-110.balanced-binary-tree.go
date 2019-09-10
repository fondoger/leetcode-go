package main

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == -1 {
		return -1
	}
	right := helper(root.Right)
	if right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
