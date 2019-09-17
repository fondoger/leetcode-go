package main

/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var best int

func diameterOfBinaryTree(root *TreeNode) int {
	best = 0
	getDepth(root)
	return best
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left+right > best {
		best = left + right
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
