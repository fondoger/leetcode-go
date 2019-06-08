package main

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findNode(dummyHead *ListNode, x int) *ListNode {
	for dummyHead.Next != nil {
		if dummyHead.Next.Val >= x {
			return dummyHead
		}
		dummyHead = dummyHead.Next
	}
	return nil
}

func partition(head *ListNode, x int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	preTarget := findNode(dummy, x)
	if preTarget == nil {
		return head
	}
	preCurrent := preTarget.Next
	for preCurrent.Next != nil {
		if preCurrent.Next.Val < x {
			tmp := preCurrent.Next
			preCurrent.Next = preCurrent.Next.Next
			tmp.Next = preTarget.Next
			preTarget.Next = tmp
			preTarget = tmp
		} else {
			preCurrent = preCurrent.Next
		}
	}
	return dummy.Next
}
