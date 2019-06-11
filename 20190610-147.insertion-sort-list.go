package main

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		// find a location to be inserted
		cur := dummy
		for cur.Next != nil && cur.Next.Val <= head.Val {
			cur = cur.Next
		}
		next := head.Next
		head.Next = cur.Next
		cur.Next = head
		head = next
	}
	return dummy.Next
}
