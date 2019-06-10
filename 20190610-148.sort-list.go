package main

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Solution 1：归并排序
func mergeLists(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	cur.Next = nil
	return dummy.Next
}

// Solution 1: 归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	head2 := slow.Next
	slow.Next = nil
	return mergeLists(sortList(head), sortList(head2))
}
