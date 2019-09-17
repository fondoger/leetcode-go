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
var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	helper(root, false)
	return sum
}
func helper(root *TreeNode, isLeft bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && isLeft {
		sum += root.Val
	}
	helper(root.Left, true)
	helper(root.Right, false)
}
