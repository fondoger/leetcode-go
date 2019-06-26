package main

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	res := 0
	helper129(&res, 0, root)
	return res
}

func helper129(res *int, tmp int, root *TreeNode) {
	if root == nil {
		return
	}
	tmp = tmp*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += tmp
		return
	}
	helper129(res, tmp, root.Left)
	helper129(res, tmp, root.Right)
}
