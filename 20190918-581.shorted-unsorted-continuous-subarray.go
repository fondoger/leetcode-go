package main

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// Solution 2:
func findUnsortedSubarray(nums []int) int {
	// Pass 1: from left to right, find the last element which is smaller
	// than its left side, mark it as end
	var end int
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < max {
			end = i
		}
	}

	// Pass 2: from right to left, find the last element which is bigger
	// than its right side, mark it as start
	var begin int
	var min = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > min {
			begin = i
		}
	}

	if begin == end {
		return 0
	}

	return end - begin + 1
}

// Solution 1: My Stupid Solution: Using Stack
func findUnsortedSubarray_Solution_1(nums []int) int {
	var s Stack
	var hasFirst bool
	var res int
	var max int

	for i, num := range nums {
		if i == 0 || num >= max {
			if !hasFirst {
				s.push(i)
			}
			max = num
		} else {
			for !s.isEmpty() && num < nums[s.top()] {
				s.pop()
			}
			if s.isEmpty() {
				res = i - (-1)
			} else {
				res = i - s.top()
			}
			hasFirst = true
		}
	}
	return res
}

type Stack []int

func (s Stack) top() int      { return s[len(s)-1] }
func (s Stack) isEmpty() bool { return len(s) == 0 }
func (s *Stack) push(t int)   { *s = append(*s, t) }
func (s *Stack) pop() int {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
