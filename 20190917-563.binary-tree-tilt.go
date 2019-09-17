package main

/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func findTilt(root *TreeNode) int {
	res = 0
	sum(root)
	return res
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := sum(root.Left)
	r := sum(root.Right)
	res += abs(l - r)
	return l + r + root.Val
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
