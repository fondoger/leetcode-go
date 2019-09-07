package main

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}
	return helper(1, n)
}

func helper(min, max int) []*TreeNode {
	var res []*TreeNode
	if min > max {
		res = append(res, nil)
		return res
	}
	for i := min; i <= max; i++ {
		left := helper(min, i-1)
		right := helper(i+1, max)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				res = append(res, root)
			}
		}
	}
	return res
}
