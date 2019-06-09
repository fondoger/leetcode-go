package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// Step 1: find middle in the list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	dummy := new(ListNode)
	dummy.Next = slow.Next
	slow.Next = nil
	// Step 2: reverse the half behind
	cur := dummy.Next
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = dummy.Next
		dummy.Next = tmp
	}
	// Step 3: merge two lists
	cur = head
	for dummy.Next != nil {
		tmp := dummy.Next
		dummy.Next = dummy.Next.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
}
