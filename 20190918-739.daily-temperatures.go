package main

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
func dailyTemperatures(T []int) []int {
	var s Stack
	res := make([]int, len(T))
	for i, t := range T {
		for !s.isEmpty() && T[s.top()] < t {
			idx := s.pop()
			res[idx] = i - idx
		}
		s.push(i)
	}
	return res
}

type Stack []int

func (s Stack) isEmpty() bool { return len(s) == 0 }
func (s Stack) top() int      { return s[len(s)-1] }
func (s *Stack) push(t int)   { *s = append(*s, t) }
func (s *Stack) pop() int {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}
