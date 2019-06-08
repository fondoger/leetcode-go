package main

func main() {
	{
		list1 := []int{1, 3, 5, 7, 9}
		list2 := []int{1}
		list3 := []int{1, 3, 5, 7}
		l1 := buildLinkedList(list1)
		l2 := buildLinkedList(list2)
		l3 := buildLinkedList(list3)
		printLinkedList(l1)
		printLinkedList(swapPairs(l1))
		printLinkedList(l2)
		printLinkedList(swapPairs(l2))
		printLinkedList(l3)
		printLinkedList(swapPairs(l3))
	}
}
