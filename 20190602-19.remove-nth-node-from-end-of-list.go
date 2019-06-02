/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
