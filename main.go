package main

func main() {
	{
		list1 := []int{1, 3, 5, 7, 9}
		list2 := []int{3, 4, 5, 6, 7}
		list3 := []int{2, 4, 6, 7, 8, 10}
		l1 := buildLinkedList(list1)
		l2 := buildLinkedList(list2)
		l3 := buildLinkedList(list3)
		printLinkedList(l1)
		printLinkedList(l2)
		printLinkedList(l3)
		listHead := mergeKLists([]*ListNode{l1, l2, l3})
		printLinkedList(listHead)
	}
}
