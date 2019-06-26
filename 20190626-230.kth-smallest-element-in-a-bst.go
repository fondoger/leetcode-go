package main

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	count := 0
	stack := make(Stack230, 0)
	cur := root
	for cur != nil || !stack.isEmpty() {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}
		node := stack.pop()
		count++
		if count == k {
			return node.Val
		}
		cur = node.Right
	}
	return -1
}

// Stack230 : a simple stack implementation
type Stack230 []*TreeNode

func (s *Stack230) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack230) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack230) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
