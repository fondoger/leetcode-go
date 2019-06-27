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

/*
 感觉这题的解法有点取巧的意思，元素一旦重复，就不适用了
*/
func deleteNod450(root *TreeNode, key int) *TreeNode {
	for root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNod450(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNod450(root.Left, key)
		return root
	} else { // root.Val == key
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else { // root.Left != nil && root.Right != nil
			rightMin := findMin(root.Right)
			root.Val = rightMin.Val
			root.Right = deleteNod450(root.Right, root.Val)
			return root
		}
	}
}

// 前置条件：root != nil
func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
