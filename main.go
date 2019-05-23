package main

func main() {
	nums := []int{
		3,
		4,
		9,
		58,
		1994,
	}
	for _, num := range nums {
		print(num, " >>> ", intToRoman(num), "\n")
	}
}
