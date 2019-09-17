package main

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	reverseInorder(root)
	return root
}

func reverseInorder(root *TreeNode) {
	if root == nil {
		return
	}
	reverseInorder(root.Right)
	root.Val += sum
	sum = root.Val
	reverseInorder(root.Left)
}
