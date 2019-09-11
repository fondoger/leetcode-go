package main

/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func prepend(s []int, t int) []int {
	s = append(s, 0)
	copy(s[1:], s)
	s[0] = t
	return s
}

// 用前序遍历的方式作弊
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack Stack
	stack.push(root)
	for !stack.isEmpty() {
		node := stack.pop()
		res = prepend(res, node.Val)
		if node.Left != nil {
			stack.push(node.Left)
		}
		if node.Right != nil {
			stack.push(node.Right)
		}
	}
	return res
}

func postorderTraversal_Recursively(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}
func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, res)
	traverse(root.Right, res)
	*res = append(*res, root.Val)
}

type Stack []*TreeNode

func (s Stack) isEmpty() bool     { return len(s) == 0 }
func (s *Stack) push(n *TreeNode) { *s = append(*s, n) }
func (s *Stack) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
