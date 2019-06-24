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
	if root == nil {
		return res
	}
	row := make([]int, 0)
	queue := make(Queue102, 0)
	queue.offer(root)
	queue.offer(nil)
	for !queue.isEmpty() {
		node := queue.poll()
		if node == nil {
			copied := make([]int, len(row))
			copy(copied, row)
			row = row[:0]
			res = append(res, copied)
			if !queue.isEmpty() {
				queue.offer(nil)
			}
		} else {
			row = append(row, node.Val)
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
		}
	}
	return res
}

type Queue102 []*TreeNode

func (q *Queue102) isEmpty() bool     { return len(*q) == 0 }
func (q *Queue102) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue102) poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
