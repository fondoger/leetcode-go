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

/*

	遇到二叉树，一律优先递归完成
	递归可以将很多复杂的问题简单话

	// 这题还可以加上一个cache, 防止重复计算

*/

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum1 := root.Val
	if root.Left != nil {
		sum1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		sum1 += rob(root.Right.Left) + rob(root.Right.Right)
	}
	sum2 := rob(root.Left) + rob(root.Right)
	return max337(sum1, sum2)
}

func max337(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
