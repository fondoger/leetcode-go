package main

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

func qsort(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	mid := low + (high-low)/2
	nums[low], nums[mid] = nums[mid], nums[low] // 在i处挖坑
	pivot := nums[low]
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
	qsort(nums, low, i-1)
	qsort(nums, i+1, high)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 10)
	qsort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		low, high := i+1, len(nums)-1
		for low < high {
			if nums[low]+nums[high] < target {
				low++
			} else if nums[low]+nums[high] > target {
				high--
			} else {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low] == nums[low-1] {
					low++
				}
				for low < high && nums[high] == nums[high+1] {
					high--
				}
			}
		}
	}
	return res
}
