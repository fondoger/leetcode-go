package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var list []int
	queue := NewQueue()
	if root != nil {
		queue.offer(root)
		queue.offer(nil)
	}
	for !queue.isEmpty() {
		t := queue.poll()
		if t != nil {
			list = append(list, t.Val)
			if t.Left != nil {
				queue.offer(t.Left)
			}
			if t.Right != nil {
				queue.offer(t.Right)
			}
		} else {
			res = append(res, list)
			list = make([]int, 0)
			if !queue.isEmpty() {
				queue.offer(nil)
			}
		}
	}
	return res
}

type Queue []*TreeNode

func NewQueue() Queue              { return make(Queue, 0) }
func (q Queue) isEmpty() bool      { return len(q) == 0 }
func (q Queue) peek() *TreeNode    { return q[0] }
func (q *Queue) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue) poll() *TreeNode {
	t := (*q)[0]
	*q = (*q)[1:]
	return t
}
