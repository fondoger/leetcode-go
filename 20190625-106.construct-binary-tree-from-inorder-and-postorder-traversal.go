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
func buildTree_REMOVE_THIS(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{postorder[n-1], nil, nil}
	idx := 0
	for inorder[idx] != postorder[n-1] {
		idx++
	}
	root.Left = buildTree_REMOVE_THIS(inorder[:idx], postorder[:idx])
	root.Right = buildTree_REMOVE_THIS(inorder[idx+1:], postorder[idx:n-1])
	return root
}
