package main

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

/*
	主要思路：扫描两边
	从左至右扫描时，将每个坑都填充为左侧所有元素之乘积
	从右至左扫描时，将每个坑都乘上右侧所有元素之乘积
*/
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	left := 1
	for i := 0; i < len(res); i++ {
		res[i] = left // 注意这里不是用的*=符号
		left *= nums[i]
	}
	right := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
