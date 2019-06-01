package main

func main() {
	workers := [][]int{{0, 0}, {1, 1}, {2, 0}}
	bikes := [][]int{{1, 0}, {2, 2}, {2, 1}}
	res := assignBikes(workers, bikes)
	println(res)

	workers2 := [][]int{{0, 0}, {2, 1}}
	bikes2 := [][]int{{1, 2}, {3, 3}}
	res2 := assignBikes(workers2, bikes2)
	println(res2)
}
