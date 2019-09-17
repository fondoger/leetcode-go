package main

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
var original int

func pathSum(root *TreeNode, sum int) int {
	res = 0
	original = sum
	helper(root, sum, true)
	return res
}
func helper(root *TreeNode, sum int, canSplit bool) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		res++
	}
	helper(root.Left, sum, true)
	helper(root.Right, sum, true)
	if canSplit {
		helper(root.Left, original, false) // starts from A
		helper(root.Right, original, false)
	}
}
