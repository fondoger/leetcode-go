package main

/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack Stack
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	cur := root
	for cur != nil {
		it.stack.push(cur)
		cur = cur.Left
	}
	return it
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.pop()
	cur := node.Right
	for cur != nil {
		this.stack.push(cur)
		cur = cur.Left
	}
	return node.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.stack.isEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type Stack []*TreeNode

func (s Stack) isEmpty() bool     { return len(s) == 0 }
func (s *Stack) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack) pop() *TreeNode {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
