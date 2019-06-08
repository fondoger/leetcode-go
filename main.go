package main

func main() {
	{
		list1 := []int{1, 3, 5, 7, 9}
		l1 := buildLinkedList(list1)
		printLinkedList(l1)
		printLinkedList(reverseKGroup(l1, 3))
	}
}
