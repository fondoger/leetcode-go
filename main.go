package main

import "reflect"

func main() {
	nums := []int{
		0,
		12321,
		-12321,
		111,
	}
	print(reflect.TypeOf('A').String())
	for _, num := range nums {
		print(num, " >>> ", isPalindrome(num), "\n")
	}
}
