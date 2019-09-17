/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxList []int

func largestValues(root *TreeNode) []int {
	maxList = nil
	traverse(root, 0)
	return maxList
}

func traverse(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(maxList) <= depth {
		maxList = append(maxList, root.Val)
	} else if root.Val > maxList[depth] {
		maxList[depth] = root.Val
	}
	traverse(root.Left, depth+1)
	traverse(root.Right, depth+1)
}

