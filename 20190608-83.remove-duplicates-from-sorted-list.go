package main

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates_REMOVE_THIS(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Next != nil && cur.Next.Next.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
