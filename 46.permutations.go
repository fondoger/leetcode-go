package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
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
