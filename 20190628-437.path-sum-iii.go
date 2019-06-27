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
func pathSum437(root *TreeNode, sum int) int {
	res := 0
	helper437(&res, root, sum, sum, false)
	return res
}

func helper437(res *int, root *TreeNode, remain int, sum int, walked bool) {
	if root == nil {
		return
	}
	remain -= root.Val
	if remain == 0 {
		*res += 1
	}
	if !walked {
		helper437(res, root.Left, sum, sum, false)
	}
	helper437(res, root.Left, remain, sum, true)
	helper437(res, root.Right, remain, sum, true)
	if !walked {
		helper437(res, root.Right, sum, sum, false)
	}
}
