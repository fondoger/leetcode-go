package main

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	left, right := root.Left, root.Left // left和right都从左子树开始
	for right != nil {
		left = left.Left
		right = right.Right
		count = count << 1
	}
	if left == nil { // 左子树是完美二叉树
		// 此时 count 代表根结点+左子树的结点数目
		return count + countNodes(root.Right)
	} else { // 左子树不是完美二叉树，则右子树必为完美二叉树
		// 此时 count 代表根结点+右子树的结点数目
		return count + countNodes(root.Left)
	}
}
