package main

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

/*
这道题花了我一整天Debug，最后在StackOverflow上解决了问题
具体请看: https://stackoverflow.com/questions/56649138
*/

func permuteUnique(nums []int) [][]int {
	qsort_47(nums, 0, len(nums)-1)
	res := make([][]int, 0, len(nums))
	helper_47(&res, nums, 0)
	return res
}

func helper_47(res *[][]int, nums []int, start int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	nums = copied
	if start == len(nums)-1 {
		*res = append(*res, nums)
		return
	}
	for i := start; i < len(nums); i++ {
		if start == i || nums[start] != nums[i] {
			nums[i], nums[start] = nums[start], nums[i] // 交换元素
			helper_47(res, nums, start+1)
			// Note: 注意这里不能交换回去
		}
	}
}

func qsort_47(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j, pivot := low, high, nums[low]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	qsort_47(nums, low, i-1)
	qsort_47(nums, i+1, high)
}
