package main

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
// 关键是存放下标
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	deque := make(Deque239, 0)
	for i, num := range nums {
		if !deque.isEmpty() && deque.peekFirst() <= i-k {
			deque.pollFirst()
		}
		for !deque.isEmpty() && nums[deque.peekLast()] <= num {
			deque.pollLast()
		}
		deque.offerLast(i)
		if i >= k-1 {
			res = append(res, nums[deque.peekFirst()])
		}
	}
	return res
}

type Deque239 []int

func (q Deque239) peekFirst() int {
	return q[0]
}

func (q Deque239) isEmpty() bool {
	return len(q) == 0
}

func (q *Deque239) pollFirst() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Deque239) offerLast(val int) {
	*q = append(*q, val)
}

func (q Deque239) peekLast() int {
	return q[len(q)-1]
}

func (q *Deque239) pollLast() int {
	val := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return val
}
