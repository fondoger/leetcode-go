package main

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Solution 1: 思路简单的办法
func detectCycle(head *ListNode) *ListNode {
	// Step 1: 判定是否有环
	slow, fast := head, head
	hasCircle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCircle = true
			break
		}
	}
	if !hasCircle {
		return nil // 无环
	}
	// Step 2: 计算环大小
	count := 1
	fast = fast.Next
	for fast != slow {
		fast = fast.Next
		count++
	}
	// Step 3: 维护一个大小为k的双指针窗口，不停向前移动，当指针首次相交时，交点即为环
	slow, fast = head, head
	for i := 0; i < count; i++ {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
