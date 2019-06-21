package main

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

func threeSum(nums []int) [][]int {
	qsort_15(nums, 0, len(nums)-1)
	res := make([][]int, 0, 10)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[low] + nums[high] + nums[i]
			if sum < 0 {
				low++
			} else if sum > 0 {
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

func qsort_15(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j, pivot := low, high, nums[(low+high)/2]
	for i < j {
		for nums[j] > pivot {
			j--
		}
		for nums[i] < pivot {
			i++
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	qsort_15(nums, low, j)
	qsort_15(nums, i, high)
}
