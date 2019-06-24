package main

import "math"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper98(root, math.MinInt64, math.MaxInt64)
}

func helper98(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val > min && root.Val < max {
		return helper98(root.Left, min, root.Val) &&
			helper98(root.Right, root.Val, max)
	}
	return false
}
