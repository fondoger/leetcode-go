package main

func main() {
	{
		list := []int{1, 2, 3, 4, 5, 6}
		head := buildLinkedList(list)
		printLinkedList(head)
		newHead := removeNthFromEnd(head, 1)
		printLinkedList(newHead)

	}
	{
		list := []int{7, 8, 9, 10, 11}
		head := buildLinkedList(list)
		printLinkedList(head)
		newHead := removeNthFromEnd(head, 5)
		printLinkedList(newHead)
	}
}
