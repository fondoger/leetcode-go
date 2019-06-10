package main

/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Solution 2: Correct Solution
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// Solution 1: Stupid Solution(依次取后一个结点的值)
func deleteNode_Stupid_Solution(node *ListNode) {

	for node.Next.Next != nil {
		node.Val, node = node.Next.Val, node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil

}
