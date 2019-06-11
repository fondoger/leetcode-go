package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head
	slow, fast := dummy, dummy
	for i := 0; i < m-1; i++ {
		slow = slow.Next
		fast = fast.Next
	}
	fast = fast.Next
	for i := m; i < n; i++ {
		tmp := fast.Next
		fast.Next = fast.Next.Next
		tmp.Next = slow.Next
		slow.Next = tmp
	}
	return dummy.Next
}
