package main

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	idx := 0
	for inorder[idx] != preorder[0] {
		idx++
	}
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	return root
}
