package main

import (
	"strconv"
	"strings"
)

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
	var res []string
	var path []string
	helper(root, &path, &res)
	return res
}

func helper(root *TreeNode, path, res *[]string) {
	if root == nil {
		return
	}
	*path = append(*path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(*path, "->"))
		*path = (*path)[:len(*path)-1]
		return
	}
	helper(root.Left, path, res)
	helper(root.Right, path, res)
	*path = (*path)[:len(*path)-1]
}
