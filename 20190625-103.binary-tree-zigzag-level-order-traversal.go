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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make(Queue103, 0)
	queue.offer(root)
	level := 0
	for !queue.isEmpty() {
		level++
		size := queue.size()
		row := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.poll()
			// calculate index on the row
			var index int
			if level%2 == 1 {
				index = i
			} else {
				index = size - i - 1
			}
			row[index] = node.Val
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
		}
		res = append(res, row)
	}
	return res
}

// Solution 1: 两个栈交替
func zigzagLevelOrder_Using_Two_Stacks(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack1 := make(Stack103, 0)
	stack2 := make(Stack103, 0)
	curStack := &stack1
	nextStack := &stack2
	curStack.push(root)
	for !curStack.isEmpty() {
		path := make([]int, 0)
		for !curStack.isEmpty() {
			node := curStack.pop()
			path = append(path, node.Val)
			if curStack == &stack1 {
				if node.Left != nil {
					nextStack.push(node.Left)
				}
				if node.Right != nil {
					nextStack.push(node.Right)
				}
			} else {
				if node.Right != nil {
					nextStack.push(node.Right)
				}
				if node.Left != nil {
					nextStack.push(node.Left)
				}
			}
		}
		res = append(res, path)
		curStack, nextStack = nextStack, curStack
	}
	return res
}

// Stack103 : simple stack
type Stack103 []*TreeNode

func (s *Stack103) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack103) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack103) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// Queue103 : simple queue
type Queue103 []*TreeNode

func (q *Queue103) size() int         { return len(*q) }
func (q *Queue103) isEmpty() bool     { return len(*q) == 0 }
func (q *Queue103) offer(t *TreeNode) { *q = append(*q, t) }
func (q *Queue103) poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
