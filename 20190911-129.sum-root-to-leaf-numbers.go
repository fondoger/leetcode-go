package main

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var res int
	helper(root, 0, &res)
	return res
}

func helper(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
	}
	helper(root.Left, sum, res)
	helper(root.Right, sum, res)
}
