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

// Solution 2: 迭代解法
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack Stack
	stack.push(root)
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

// Solution 1: 递归解法
func preorderTraversal_Recursively(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	traverse(root.Left, res)
	traverse(root.Right, res)
}

type Stack []*TreeNode

func (s Stack) isEmpty() bool     { return len(s) == 0 }
func (s *Stack) push(n *TreeNode) { *s = append(*s, n) }
func (s *Stack) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
