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
	res := make([][]int, 0)
	path := make([]int, 0)
	helper113(&res, root, sum, &path)
	return res
}

func helper113(res *[][]int, root *TreeNode, sum int, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		copied := make([]int, len(*path))
		copy(copied, *path)
		*res = append(*res, copied)
	} else {
		helper113(res, root.Left, sum, path)
		helper113(res, root.Right, sum, path)
	}
	*path = (*path)[:len(*path)-1]
}
