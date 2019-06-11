package main

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		haveDuplicate := false
		for cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			cur.Next = cur.Next.Next
			haveDuplicate = true
		}
		if haveDuplicate {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
