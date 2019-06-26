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

// 关键：前序遍历
func flatten(root *TreeNode) {
	helper114(root)
}

// flattern并返回尾部结点
func helper114(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil {
		return helper114(root.Right)
	} else if root.Right == nil {
		root.Right, root.Left = root.Left, nil
		return helper114(root.Right)
	}
	leftRoot := root.Left
	root.Left = nil
	leftTail := helper114(leftRoot)
	root.Right, leftTail.Right = leftRoot, root.Right
	return helper114(leftTail.Right)
}
