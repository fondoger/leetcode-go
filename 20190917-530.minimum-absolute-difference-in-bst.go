package main

import "math"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var prev *TreeNode
var minDiff int

func getMinimumDifference(root *TreeNode) int {
	prev = nil
	minDiff = math.MaxInt64
	inorder(root)
	return minDiff
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	if prev != nil && abs(prev.Val-root.Val) < minDiff {
		minDiff = abs(prev.Val - root.Val)
	}
	prev = root
	inorder(root.Right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
