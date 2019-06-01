package main

import "math"

func abs(t int) int {
	if t >= 0 {
		return t
	} else {
		return -t
	}
}

func helper(workers [][]int, bikes [][]int, used []int, path []int, cur int, res *int) {
	if cur >= len(workers) {
		sum := 0
		for w, b := range path {
			sum += abs(workers[w][0]-bikes[b][0]) + abs(workers[w][1]-bikes[b][1])
		}
		if sum < *res {
			*res = sum
		}
		return
	}
	for i := 0; i < len(bikes); i++ {
		if used[i] == 0 {
			used[i] = 1
			path[cur] = i
			helper(workers, bikes, used, path, cur+1, res)
			used[i] = 0
		}
	}
}

func assignBikes(workers [][]int, bikes [][]int) int {
	if len(workers) == 0 {
		return 0
	}
	path := make([]int, len(workers))
	res := math.MaxInt32
	used := make([]int, len(bikes))
	helper(workers, bikes, used, path, 0, &res)
	return res
}
