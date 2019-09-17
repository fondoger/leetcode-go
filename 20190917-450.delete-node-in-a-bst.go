package main

/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}
	rightMin := findMin(root.Right)
	root.Val = rightMin.Val
	root.Right = deleteNode(root.Right, rightMin.Val)
	return root
}

// 注意：找二叉树的最小值不必递归，迭代就行了
func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
