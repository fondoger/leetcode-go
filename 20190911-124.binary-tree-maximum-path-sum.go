package main

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
	var res = root.Val
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, res)
	right := helper(root.Right, res)
	sum := root.Val
	if left > 0 {
		sum += left
	}
	if right > 0 {
		sum += right
	}
	if sum > *res {
		*res = sum
	}
	return root.Val + max(0, max(left, right))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
