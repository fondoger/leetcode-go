package main

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
func merge(intervals [][]int) [][]int {
	qsort(intervals, 0, len(intervals)-1)
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return res
	}
	begin, end := intervals[0][0], intervals[0][1]
	for i := 1; i <= len(intervals); i++ {
		if i == len(intervals) || intervals[i][0] > end {
			res = append(res, []int{begin, end})
			if i != len(intervals) {
				begin, end = intervals[i][0], intervals[i][1]
			}
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}
	return res
}

func qsort(nums [][]int, low, high int) {
	if low >= high {
		return
	}
	i, j, pivot := low, high, nums[low]
	for i < j {
		for i < j && nums[j][0] >= pivot[0] {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i][0] <= pivot[0] {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	qsort(nums, low, i-1)
	qsort(nums, i+1, high)
}
