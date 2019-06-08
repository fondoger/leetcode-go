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

import (
	"container/heap"
	"sort"
)

type MinHeap []*ListNode              // 提供一个切片用作堆元素的容器
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] } // 喜欢这个语法糖
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/* Solution 3: 构造最小堆 时间复杂度 */
func mergeKLists(lists []*ListNode) *ListNode {
	priorityQueue := make(MinHeap, 0, len(lists))
	for _, list := range lists {
		if list != nil {
			priorityQueue = append(priorityQueue, list)
		}
	}
	heap.Init(&priorityQueue) // 为什么用指针呢？因为会修改切片的指向
	dummy := new(ListNode)
	cur := dummy
	for priorityQueue.Len() > 0 {
		//head := priorityQueue.Pop().(*ListNode) // 错误！
		head := heap.Pop(&priorityQueue).(*ListNode)
		cur.Next = head
		cur = cur.Next
		if head.Next != nil {
			heap.Push(&priorityQueue, head.Next)
		}
	}
	return dummy.Next
}

/* Solution 2: 全部放入列表排序 时间复杂度O(nk*log nk) */
func mergeKLists_2(lists []*ListNode) *ListNode {
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
