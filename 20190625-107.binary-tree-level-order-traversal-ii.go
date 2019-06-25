package main

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func prependRow(x [][]int, y []int) [][]int {
	x = append(x, nil)
	copy(x[1:], x)
	x[0] = y
	return x
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make(Queue107, 0)
	queue.offer(root)
	for !queue.isEmpty() {
		size := queue.size()
		row := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.poll()
			row[i] = node.Val
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
		}
		res = prependRow(res, row)
	}
	return res
}

type Queue107 []*TreeNode

func (q *Queue107) size() int         { return len(*q) }
func (q *Queue107) isEmpty() bool     { return len(*q) == 0 }
func (q *Queue107) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue107) poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
