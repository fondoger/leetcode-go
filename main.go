package main

func main() {
	{
		list1 := []int{1, 4, 3, 2, 5, 2}
		l1 := buildLinkedList(list1)
		printLinkedList(l1)
		printLinkedList(partition(l1, 3))
	}
}
