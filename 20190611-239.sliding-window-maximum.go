package main

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	deque := make([]int, 0, len(nums))
	for i, val := range nums {
		if len(deque) > 0 && PeekFirst(&deque) <= i-k {
			PollFirst(&deque)
		}
		for len(deque) > 0 && nums[PeekLast(&deque)] < val {
			PollLast(&deque)
		}
		OfferLast(&deque, i)
		if i >= k-1 {
			res = append(res, nums[PeekFirst(&deque)])
		}
	}
	return res
}

// PeekLast : 模拟双端队列
func PeekLast(nums *[]int) int {
	return (*nums)[len(*nums)-1]
}

// PollLast : 模拟双端队列
func PollLast(nums *[]int) int {
	res := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	return res
}

// PeekFirst : 模拟双端队列
func PeekFirst(nums *[]int) int {
	return (*nums)[0]
}

// PollFirst : 模拟双端队列
func PollFirst(nums *[]int) int {
	res := (*nums)[0]
	*nums = (*nums)[1:]
	return res
}

// OfferLast : 模拟双端队列
func OfferLast(nums *[]int, val int) {
	*nums = append(*nums, val)
}
