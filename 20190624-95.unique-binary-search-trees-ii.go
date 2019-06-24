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
		return []*TreeNode{}
	}
	return helper95(1, n)
}

func helper95(low, high int) []*TreeNode {
	res := []*TreeNode{}
	if low > high {
		res := append(res, nil)
		return res
	}
	for rootVal := low; rootVal <= high; rootVal++ {
		leftChildren := helper95(low, rootVal-1)
		rightChildren := helper95(rootVal+1, high)
		for _, left := range leftChildren {
			for _, right := range rightChildren {
				root := &TreeNode{rootVal, nil, nil}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
