package main

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var hashMap map[*TreeNode]int

func rob(root *TreeNode) int {
	hashMap = make(map[*TreeNode]int, 0)
	return helper(root)
}
func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if val, ok := hashMap[root]; ok {
		return val
	}
	var includeRoot, excludeRoot int
	excludeRoot = helper(root.Left) + helper(root.Right)
	includeRoot = root.Val
	if root.Left != nil {
		includeRoot += helper(root.Left.Left) + helper(root.Left.Right)
	}
	if root.Right != nil {
		includeRoot += helper(root.Right.Left) + helper(root.Right.Right)
	}
	max := includeRoot
	if includeRoot < excludeRoot {
		max = excludeRoot
	}
	hashMap[root] = max
	return max
}
