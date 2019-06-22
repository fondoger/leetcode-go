package main

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
func longestValidParentheses(s string) int {
	stack := make(Stack32, 0)
	res := 0
	stack.push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.push(i)
		} else {
			stack.pop()
			if !stack.isEmpty() {
				res = max32(res, i-stack.peek())
			} else {
				stack.push(i)
			}
		}
	}
	return res
}

func max32(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Stack32 []int

func (s *Stack32) push(i int) {
	*s = append(*s, i)
}
func (s *Stack32) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *Stack32) peek() int {
	return (*s)[len(*s)-1]
}
func (s *Stack32) isEmpty() bool {
	return len(*s) == 0
}
