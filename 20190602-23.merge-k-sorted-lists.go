/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for true {
		curBestIdx := -1
		for i, list := range lists {
			if list != nil && (curBestIdx == -1 || list.Val < lists[curBestIdx].Val) {
				curBestIdx = i
			}
		}
		if curBestIdx != -1 {
			cur.Next = lists[curBestIdx]
			lists[curBestIdx] = lists[curBestIdx].Next
			cur = cur.Next
		} else {
			break
		}
	}
	return dummy.Next
}
