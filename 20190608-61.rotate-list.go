package main

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getSize(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	size := getSize(head)
	k = k % size
	if k == 0 {
		return head
	}
	// Two pointers
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}
