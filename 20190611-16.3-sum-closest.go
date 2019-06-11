package main

import "math"

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// 快速排序交换法
func qsort2(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j, pivot := low, high, nums[low+(high-low)>>2]
	for i < j {
		for nums[i] < pivot { // 与填充法的区别：不考虑 i 递增越界，不取等号
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i <= j { // 注意这里是等号
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	qsort2(nums, low, j)  // low 对应 j
	qsort2(nums, i, high) // high 对应 i
}

// 就是一个普通的取绝对值函数，用如此丑陋的方式仅仅是防止在package main中重复定义
func abs_16(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	bestSum := 0
	minDiff := math.MaxInt32
	qsort2(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high := i+1, len(nums)-1
		diff := nums[i] - target
		for low < high {
			if abs_16(nums[low]+nums[high]+diff) < minDiff {
				minDiff = abs_16(nums[low] + nums[high] + diff)
				bestSum = nums[low] + nums[high] + nums[i]
			}
			if nums[low]+nums[high]+diff < 0 {
				low++
			} else if (nums[high] + nums[high] + diff) > 0 {
				high--
			} else {
				return target
			}
		}
	}
	return bestSum
}
