package main

import "reflect"

func main() {
	nums := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
	}
	print(reflect.TypeOf('A').String())
	for _, num := range nums {
		print(num, " >>> ", myAtoi(num), "\n")
	}
}
