package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// Solution 2: 单调栈(十分类似单调队列)
// 只需确保栈中元素是单调递增就好
func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make(stack_84, 0, len(heights))
	heights = append(heights, -1) // 在尾部加一个元素,确保最后栈中元素被全部弹出
	for i := 0; i < len(heights); i++ {
		for !stack.isEmpty() && heights[stack.peek()] > heights[i] {
			h := heights[stack.pop()]
			var w int
			if stack.isEmpty() {
				w = i
			} else {
				w = i - stack.peek() - 1
			}
			if h*w > maxArea {
				maxArea = h * w
			}
		}
		stack.push(i)
	}
	return maxArea
}

// Solution 1: 预处理，遍历每个结点，
// 重点在 p=fromLeft[p] 代替 p--, 将时间复杂度降低一个级别
// 巧妙至极
func largestRectangleArea_1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	fromLeft := make([]int, len(heights))
	fromRight := make([]int, len(heights))
	fromLeft[0] = -1
	fromRight[len(heights)-1] = len(heights)

	for i := 1; i < len(heights); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = fromLeft[p]
		}
		fromLeft[i] = p
	}
	for i := len(heights) - 2; i >= 0; i-- {
		p := i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = fromRight[p]
		}
		fromRight[i] = p
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		area := heights[i] * (fromRight[i] - fromLeft[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 使用切片实现的简单栈
type stack_84 []int

func (s *stack_84) isEmpty() bool {
	return len(*s) == 0
}
func (s *stack_84) peek() int {
	return (*s)[len(*s)-1]
}
func (s *stack_84) push(v int) {
	*s = append(*s, v)
}
func (s *stack_84) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
