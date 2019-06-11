package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome_REMOVE_THIS(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// Step 1: find middle of the list, and partition it
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// Step 2: Reverse the list behind
	head2 := reverseLinkedList(slow.Next)
	slow.Next = nil // !important
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func reverseLinkedList(head *ListNode) *ListNode {
	var tail *ListNode = nil
	for head != nil {
		head, tail, head.Next = head.Next, head, tail
	}
	return tail
}
