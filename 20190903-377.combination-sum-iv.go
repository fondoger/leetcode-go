package main

/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// Note that different sequences are counted as different combinations.

func combinationSum4(nums []int, target int) int {
	m := make(map[int]int, 0)
	return helper377(m, nums, target)
}

func helper377(m map[int]int, nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	if val, ok := m[target]; ok {
		return val
	}
	res := 0
	for _, num := range nums {
		count := helper377(m, nums, target-num)
		res += count
	}
	m[target] = res
	return res
}
