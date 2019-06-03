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

import "sort"

/* Solution 2: 全部放入列表排序 时间复杂度O(nk*log nk) */
func mergeKLists(lists []*ListNode) *ListNode {
	array := make([]*ListNode, 0, 100)
	// step1: put all nodes into a list
	for _, list := range lists {
		for list != nil {
			array = append(array, list)
			list = list.Next
		}
	}
	// step 2: sort by quick sort
	sort.Slice(array, func(i, j int) bool {
		return array[i].Val < array[j].Val
	})
	// step 3: rebuild list
	dummy := new(ListNode)
	cur := dummy
	for _, node := range array {
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = nil
	return dummy.Next
}

/* Solution 1: 双重for循环 时间复杂度: nk * k (k个链表,每个链表平均长度为n) */
func mergeKLists_1(lists []*ListNode) *ListNode {
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
