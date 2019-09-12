package main

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var res int
	helper(root, &k, &res)
	return res
}

func helper(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}
	helper(root.Left, k, res)
	*k = *k - 1
	if *k == 0 {
		*res = root.Val
		return
	}
	helper(root.Right, k, res)
}
