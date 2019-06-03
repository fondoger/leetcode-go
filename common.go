package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入：int[]数组
// 输出：ListNode单向链表
func buildLinkedList(nums []int) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for _, val := range nums {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return dummy.Next
}

// 打印 ListNode 单向链表
func printLinkedList(head *ListNode) {
	print("List: ")
	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
	println()
}

// 调用元素默认的toString()方法
func printIntArray(arr []int) {
	print("Array: ")
	for _, item := range arr {
		print(item, " ")
	}
	println()
}
