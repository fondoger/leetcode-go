package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	m := postorder[len(postorder)-1]
	idx := 0
	for i, val := range inorder {
		if val == m {
			idx = i
		}
	}
	root := &TreeNode{m, nil, nil}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}
