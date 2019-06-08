package main

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 在不改变时间复杂度的情况下，可以增加遍历次数来简化代码
func isMoreThanKNodes(head *ListNode, k int) bool {
	count := 0
	for count < k {
		if head == nil {
			return false
		}
		count++
		head = head.Next
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for isMoreThanKNodes(cur.Next, k) {
		dummy2 := cur
		cur = cur.Next
		for count := 1; count < k; count++ {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = dummy2.Next
			dummy2.Next = tmp
		}
	}

	return dummy.Next
}
