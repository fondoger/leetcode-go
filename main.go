package main

func main() {
	if true {
		arr := [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		}
		// arr := make([][]int, 3)
		// arr[0] = []int{0, 0, 0} // slice
		// arr[1] = []int{1, 1, 0} // slice

		println(shortestPathBinaryMatrix(arr))
	}
}
