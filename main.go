package main

func main() {
	{
		list1 := []int{1, 3, 5, 7, 9}
		list2 := []int{3, 4, 5, 6, 7}
		l1 := buildLinkedList(list1)
		l2 := buildLinkedList(list2)
		printLinkedList(l1)
		printLinkedList(l2)
		listHead := mergeTwoLists(l1, l2)
		printLinkedList(listHead)
	}
}
