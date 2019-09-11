package main

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution 2: 巧妙解法
var prev *TreeNode

func flatten(root *TreeNode) {
	prev = nil
	flat(root)
}
func flat(root *TreeNode) {
	if root == nil {
		return
	}
	flat(root.Right)
	flat(root.Left)
	root.Right = prev
	root.Left = nil
	prev = root
}

// Solution 1: 直觉写法
func flatten_(root *TreeNode) {
	helper(root)
}
func helper(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil {
		return helper(root.Right)
	} else if root.Right == nil {
		root.Right, root.Left = root.Left, nil
		return helper(root.Right)
	}
	left := helper(root.Left)
	right := helper(root.Right)
	left.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return right
}
