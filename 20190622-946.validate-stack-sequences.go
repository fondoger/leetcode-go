package main

/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */

// 最关键的一点就是假设栈中无重复数字
func validateStackSequences(pushed []int, popped []int) bool {
	stack := NewStack946()
	pos := 0
	for _, val := range pushed {
		stack.push(val)
		for !stack.isEmpty() && stack.peek() == popped[pos] {
			stack.pop()
			pos++
		}
	}
	return stack.isEmpty()
}

type Node946 = int // 类型别名
type Stack946 []Node946

func NewStack946() Stack946 {
	return make(Stack946, 0)
}
func (s *Stack946) peek() Node946 {
	return (*s)[len(*s)-1]
}
func (s *Stack946) pop() Node946 {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *Stack946) push(val Node946) {
	*s = append(*s, val)
}
func (s *Stack946) isEmpty() bool {
	return len(*s) == 0
}
