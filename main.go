package main

func main() {
	if true {
		nums := [][]int{
			buildIntArray("[2,1,5]"),
			buildIntArray("[3,3,7]"),
		}

		println(carPooling(nums, 4))
	}
	if true {
		nums := [][]int{
			buildIntArray("[2,1,5]"),
			buildIntArray("[3,3,7]"),
		}

		println(carPooling(nums, 5))
	}
	if true {
		nums := [][]int{
			buildIntArray("[2,1,5]"),
			buildIntArray("[3,5,7]"),
		}

		println(carPooling(nums, 3))
	}
	if true {
		nums := [][]int{
			buildIntArray("[3,2,7]"),
			buildIntArray("[3,7,9]"),
			buildIntArray("[8,3,9]"),
		}

		println(carPooling(nums, 11))
	}
	if true {
		nums := [][]int{
			buildIntArray("[8,2,3]"),
			buildIntArray("[4,1,3]"),
			buildIntArray("[1,3,6]"),
			buildIntArray("[8,4,6]"),
			buildIntArray("[4,4,8]"),
		}

		println(carPooling(nums, 12))
	}
}
