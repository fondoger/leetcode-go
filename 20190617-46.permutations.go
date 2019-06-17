package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// 方法一：选择法 + 挑选记录数组
func permute(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	helper_46_2(&res, nums, make([]int, len(nums)), 0, make([]bool, len(nums)))
	return res
}

func helper_46_2(res *[][]int, nums, path []int, start int, used []bool) {
	if start == len(nums) {
		// copy path
		copied := make([]int, len(path))
		copy(copied, path)
		*res = append(*res, copied)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			path[start] = nums[i]
			helper_46_2(res, nums, path, start+1, used)
			used[i] = false // 关键
		}
	}
}

// Solution 1: 交换法
func permute_Solution1(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	helper_46(&res, nums, 0)
	return res
}

func helper_46(res *[][]int, nums []int, start int) {
	if start == len(nums) {
		path := make([]int, len(nums))
		copy(path, nums)
		*res = append(*res, path)
		return
	}
	for i := start; i < len(nums); i++ {
		swap_46(&nums[start], &nums[i])
		helper_46(res, nums, start+1)
		swap_46(&nums[start], &nums[i])
	}
}

func swap_46(i, j *int) {
	*i, *j = *j, *i
}
