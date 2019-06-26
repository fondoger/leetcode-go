package main

/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

func lowestCommonAncestor_PLEASE_REMOVE_THIS(root, p, q *TreeNode) *TreeNode {
	pathP := make([]*TreeNode, 0)
	pathQ := make([]*TreeNode, 0)
	pathToVal(&pathP, root, p.Val)
	pathToVal(&pathQ, root, q.Val)
	min := min236(len(pathP), len(pathQ))
	for i := 0; i < min; i++ {
		if i == min-1 {
			return pathP[min-1]
		} else if pathP[i+1] != pathQ[i+1] {
			return pathP[i]
		}
	}
	return nil
}

func pathToVal(path *[]*TreeNode, root *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root.Val == val {
		return true
	}
	if pathToVal(path, root.Left, val) {
		return true
	}
	if pathToVal(path, root.Right, val) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func min236(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
