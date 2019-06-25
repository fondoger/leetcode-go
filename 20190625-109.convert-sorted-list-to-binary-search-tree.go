package main

/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var dummy = &ListNode{9527, nil}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	dummy.Next = head
	slow, fast := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := &TreeNode{slow.Next.Val, nil, nil}
	head2 := slow.Next.Next
	slow.Next = nil
	head1 := dummy.Next
	root.Left = sortedListToBST(head1)
	root.Right = sortedListToBST(head2)
	return root
}
