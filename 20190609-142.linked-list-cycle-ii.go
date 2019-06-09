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

// Solution 2: 巧妙的办法
// 在解法1的基础上，我们可以证明：从起点到第一次相遇的点之间的窗口大小恰好是环的大小的整数倍
// 这样就免去了计算环大小和构建滑动窗口的步骤
// 证明如下：
// 		设链表由长度为 p 的非环路径和长度为 C 的环路径组成。
//	 	设在第一次相遇时，相遇点到环起点的距离为 m
// 		  s_slow = p + m + i*C
//		2*s_slow = p + m + j*C = s_fast
//	 相减得s_slow = (j - i) * C，即双方同时到达相遇点时，乌龟走的路径长度恰好是环长度的整数倍，
// 	 	因此，当乌龟第一次经过相遇点时，此时已经走了环长度的整数倍。
/// 这个证明不完善，不要看！！！
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil // 无环
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// Solution 1: 思路简单的办法
func detectCycle_1(head *ListNode) *ListNode {
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
	// Step 3: 维护一个大小为count的双指针窗口，不停向前移动，当指针首次相交时，交点即为环
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
