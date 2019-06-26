package main

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	queue := make(Queue199, 0)
	if root != nil {
		queue.offer(root)
	}
	for !queue.isEmpty() {
		size := queue.size()
		rightOne := queue.peek()
		res = append(res, rightOne.Val)
		for i := 0; i < size; i++ {
			node := queue.poll()
			if node.Right != nil {
				queue.offer(node.Right)
			}
			if node.Left != nil {
				queue.offer(node.Left)
			}
		}
	}
	return res
}

// Queue199 : a simple queue implementation
type Queue199 []*TreeNode

func (q *Queue199) size() int         { return len(*q) }
func (q *Queue199) isEmpty() bool     { return len(*q) == 0 }
func (q *Queue199) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue199) peek() *TreeNode   { return (*q)[0] }
func (q *Queue199) poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
