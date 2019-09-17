package main

/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var hashMap map[int]int
var maxCount int

func findFrequentTreeSum(root *TreeNode) []int {
	hashMap = make(map[int]int, 0)
	maxCount = 0
	postorder(root)
	var res []int
	for k, v := range hashMap {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}
func postorder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := postorder(root.Left)
	right := postorder(root.Right)
	sum := root.Val + left + right
	hashMap[sum]++
	if hashMap[sum] > maxCount {
		maxCount = hashMap[sum]
	}
	return sum
}
