package main

/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(tmp)
	return root
}
