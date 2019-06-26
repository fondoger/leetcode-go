package main

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution 2: 二叉树前序遍历：迭代解法
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make(Stack144, 0)
	if root != nil {
		stack.push(root)
	}
	for !stack.isEmpty() {
		node := stack.pop()
		res = append(res, node.Val)
		if node.Right != nil {
			stack.push(node.Right)
		}
		if node.Left != nil {
			stack.push(node.Left)
		}
	}
	return res
}

// Solution 1: 二叉树前序遍历：递归解法
func preorderTraversal_RECURSIVE_SOLUTION(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(&res, root)
	return res
}

func preorder(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(res, root.Left)
	preorder(res, root.Right)
}

// Stack144 : a simple stack implementation
type Stack144 []*TreeNode

func (s *Stack144) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack144) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack144) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
