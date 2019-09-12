package main

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	traverse(root, 0, &res)
	return res
}
func traverse(root *TreeNode, level int, res *[]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, 0)
	}
	(*res)[level] = root.Val
	traverse(root.Left, level+1, res)
	traverse(root.Right, level+1, res)
}
