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

// 解法二：非递归解法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 10)
	stack := NewStack94()
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

// 解法1：递归解法
func inorderTraversal_Recursive_Solution(root *TreeNode) []int {
	res := make([]int, 0, 10)
	helper94(&res, root)
	return res
}

func helper94(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	helper94(res, root.Left)
	*res = append(*res, root.Val)
	helper94(res, root.Right)
}

type Node94 = *TreeNode // 类型别名
type Stack94 []Node94

func NewStack94() Stack94 {
	return make(Stack94, 0)
}
func (s *Stack94) peek() Node94 {
	return (*s)[len(*s)-1]
}
func (s *Stack94) pop() Node94 {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *Stack94) push(val Node94) {
	*s = append(*s, val)
}
func (s *Stack94) isEmpty() bool {
	return len(*s) == 0
}
