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
	return depth110(root) != -1
}

func depth110(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth110(root.Left)
	if left == -1 {
		return -1
	}
	right := depth110(root.Right)
	if right == -1 {
		return -1
	}
	if abs110(left-right) > 1 {
		return -1
	}
	return max110(left, right) + 1
}
func abs110(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func max110(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
