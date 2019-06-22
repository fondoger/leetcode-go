package main

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

/*
	最长括号匹配

	主要有两种解法：
	1. （简单的）栈解法，关键在于存放下标，而非存放左括号
	2. （稍复杂的）动态规划解法

*/

// 解法2：动态规划
// 数组dp[i]存放以 i 为结尾的最长括号匹配
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	}
	res := dp[1]
	for i := 2; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else if dp[i-1] > 0 { // s[i-1]必为右括号
				preIdx := i - dp[i-1] - 1
				if preIdx >= 0 && s[preIdx] == '(' {
					dp[i] = dp[i-1] + 2
					if preIdx-1 >= 0 {
						dp[i] += dp[preIdx-1]
					}
				}
			}
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}

// 解法1：使用栈，关键在存放下标而非左括号
func longestValidParentheses_Using_Stack(s string) int {
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
