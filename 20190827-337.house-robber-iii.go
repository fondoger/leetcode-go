package main

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 树型抢劫, 递归 + memo缓存
// 添加memo很简单，略
func robTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	num1 := root.Val
	if root.Left != nil {
		num1 += robTree(root.Left.Left) + robTree(root.Left.Right)
	}
	if root.Right != nil {
		num1 += robTree(root.Right.Left) + robTree(root.Right.Right)
	}
	num2 := robTree(root.Left) + robTree(root.Right)
	return max337(num1, num2)
}

func rob_REMOVE_THIS_PLEASE(root *TreeNode) int {
	return robTree(root)
}

func max337(a, b int) int {
	if a > b {
		return a
	}
	return b
}
