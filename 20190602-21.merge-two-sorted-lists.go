/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l2.Val <= l1.Val) {
			cur.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
			cur = cur.Next
		} else {
			cur.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
			cur = cur.Next
		}
	}
	cur.Next = nil
	return dummy.Next
}
