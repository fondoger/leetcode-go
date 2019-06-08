package main

func main() {
	{
		list1 := []int{1, 2, 3, 3, 4, 4, 5}
		l1 := buildLinkedList(list1)
		printLinkedList(l1)
		printLinkedList(deleteDuplicates(l1))
	}
}
