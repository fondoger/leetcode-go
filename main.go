package main

func main() {
	{
		list1 := []int{1, 2, 3, 4, 5}
		l1 := buildLinkedList(list1)
		printLinkedList(l1)
		printLinkedList(reverseBetween(l1, 2, 4))
	}
}
