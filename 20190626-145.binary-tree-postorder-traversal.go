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

func prepend145(s []int, val int) []int {
	copied := append(s, 0)
	copy(copied[1:], s)
	copied[0] = val
	return copied
}

// Solution 2: 二叉树的后序遍历：使用栈的迭代解法, 作弊解法，逆序输出
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make(Stack145, 0)
	if root != nil {
		stack.push(root)
	}
	for !stack.isEmpty() {
		node := stack.pop()
		res = prepend145(res, node.Val)
		if node.Left != nil {
			stack.push(node.Left)
		}
		if node.Right != nil {
			stack.push(node.Right)
		}
	}
	return res
}

// Solution 1: 二叉树的后序遍历：递归解法
func postorderTraversal_RECURSIVE_SOLUTION(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(&res, root)
	return res
}

func postorder(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	postorder(res, root.Left)
	postorder(res, root.Right)
	*res = append(*res, root.Val)
}

// Stack145 : a simple stack implementation
type Stack145 []*TreeNode

func (s *Stack145) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack145) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack145) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
