package main

/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution 2: Using Queue
// 仍然按照普通层序遍历方式遍历，但插入时用trick计算插入位置
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var q Queue
	q.offer(root)
	level := 0
	for !q.isEmpty() {
		level++
		size := q.size()
		list := make([]int, size)
		for i := 0; i < size; i++ {
			n := q.poll()
			var idx int
			if level%2 == 0 {
				idx = size - i - 1
			} else {
				idx = i
			}
			list[idx] = n.Val
			if n.Left != nil {
				q.offer(n.Left)
			}
			if n.Right != nil {
				q.offer(n.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

// Solution 1: Using Stack
func zigzagLevelOrder_Using_Stack(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var stack1, stack2 Stack
	cur, another := &stack1, &stack2
	cur.push(root)
	for !cur.isEmpty() {
		path := make([]int, 0)
		for !cur.isEmpty() {
			n := cur.pop()
			path = append(path, n.Val)
			if cur == &stack1 {
				if n.Left != nil {
					another.push(n.Left)
				}
				if n.Right != nil {
					another.push(n.Right)
				}
			} else {
				if n.Right != nil {
					another.push(n.Right)
				}
				if n.Left != nil {
					another.push(n.Left)
				}
			}
		}
		res = append(res, path)
		cur, another = another, cur
	}
	return res
}

type Stack []*TreeNode

func (s *Stack) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

type Queue []*TreeNode

func (q Queue) isEmpty() bool      { return len(q) == 0 }
func (q Queue) size() int          { return len(q) }
func (q *Queue) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue) poll() *TreeNode {
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}
