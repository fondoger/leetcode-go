package main

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = tmp.Next
		tmp.Next = cur.Next
		cur.Next = tmp
		cur = tmp.Next
	}
	return dummy.Next
}
