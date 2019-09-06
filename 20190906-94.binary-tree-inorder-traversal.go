package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法2: 迭代解法
func inorderTraversal(root *TreeNode) []int {
	var res []int
	stack := NewStack()
	cur := root
	for cur != nil || !stack.isEmpty() {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}
		cur = stack.pop()
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// 解法1: 递归解法
func inorderTraversal_Recursively(root *TreeNode) []int {
	var res []int
	helper(&res, root)
	return res
}
func helper(res *[]int, root *TreeNode) {
	if root != nil {
		helper(res, root.Left)
		*res = append(*res, root.Val)
		helper(res, root.Right)
	}
}

type Stack []*TreeNode

func NewStack() *Stack         { return &Stack{} }
func (s Stack) isEmpty() bool  { return len(s) == 0 }
func (s Stack) top() *TreeNode { return (s)[len(s)-1] }
func (s *Stack) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
func (s *Stack) push(n *TreeNode) { *s = append(*s, n) }
