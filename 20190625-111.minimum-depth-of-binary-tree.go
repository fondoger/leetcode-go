package main

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	depth := 0
	queue := make(Queue111, 0)
	if root != nil {
		queue.offer(root)
	}
	for !queue.isEmpty() {
		depth++
		size := queue.size()
		for i := 0; i < size; i++ {
			node := queue.poll()
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
		}
	}
	return depth
}

// Queue111 : simple queue
type Queue111 []*TreeNode

func (q *Queue111) size() int         { return len(*q) }
func (q *Queue111) isEmpty() bool     { return len(*q) == 0 }
func (q *Queue111) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue111) poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
