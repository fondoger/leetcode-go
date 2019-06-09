package main

func main() {
	l1 := buildLinkedList("[4,2,1,3]")
	printLinkedList(l1)
	printLinkedList(insertionSortList(l1))
}
