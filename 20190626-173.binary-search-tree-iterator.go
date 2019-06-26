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
	stack *Stack173
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make(Stack173, 0)
	bst := BSTIterator{&stack}
	for root != nil {
		bst.stack.push(root)
		root = root.Left
	}
	return bst
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack.pop()
	res := node.Val
	node = node.Right
	for node != nil {
		this.stack.push(node)
		node = node.Left
	}
	return res
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

// Stack173 : a simple stack implementation
type Stack173 []*TreeNode

func (s *Stack173) isEmpty() bool    { return len(*s) == 0 }
func (s *Stack173) push(t *TreeNode) { *s = append(*s, t) }
func (s *Stack173) pop() *TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
