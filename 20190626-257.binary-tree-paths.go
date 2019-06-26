package main

import "strconv"

/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	path := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, path)
		return res
	}
	helper257(&res, path, root.Left)
	helper257(&res, path, root.Right)
	return res
}

func helper257(res *[]string, path string, root *TreeNode) {
	if root == nil {
		return
	}
	path = path + "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	helper257(res, path, root.Left)
	helper257(res, path, root.Right)
}
