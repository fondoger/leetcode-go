package main

/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	traverse(root, sum, []int{}, &res)
	return res
}

func traverse(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	sum -= root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			*res = append(*res, copySlice(path))
		}
		return
	}
	traverse(root.Left, sum, path, res)
	traverse(root.Right, sum, path, res)
}

func copySlice(s []int) []int {
	t := make([]int, len(s))
	copy(t, s)
	return t
}
