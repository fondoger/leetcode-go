package main

func main() {
	l1 := buildLinkedList("[1,2,2,1]")
	l2 := buildLinkedList("[]")
	l3 := buildLinkedList("[1]")
	l4 := buildLinkedList("[1,2,3,4,5,4,3,2,1]")
	println(isPalindrome(l1))
	println(isPalindrome(l2))
	println(isPalindrome(l3))
	println(isPalindrome(l4))
}
