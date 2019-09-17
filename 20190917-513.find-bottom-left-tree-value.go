package main

/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res, deepDepth int

func findBottomLeftValue(root *TreeNode) int {
	res, deepDepth = 0, 0
	traverse(root, 0)
	return res
}

func traverse(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth++
	if depth > deepDepth {
		res = root.Val
		deepDepth = depth
	}
	traverse(root.Left, depth)
	traverse(root.Right, depth)
}
