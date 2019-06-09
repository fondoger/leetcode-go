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

var dummy = new(ListNode)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	dummy.Next = head
	slow, fast := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = sortedListToBST(dummy.Next)
	root.Right = sortedListToBST(mid.Next)
	return root
}
