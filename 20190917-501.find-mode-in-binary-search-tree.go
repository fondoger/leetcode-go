package main

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 98. Validate Binary Search Tree 夹逼方法 或者 利用有序数组
// 99. Recover Binary Search Tree 利用有序数组
// 501.Find Mode In Binary Search Tree 利用有序数组

var prev *TreeNode
var count = 0
var maxCount = 0
var list []int

func findMode(root *TreeNode) []int {
	prev, count, maxCount, list = nil, 0, 0, nil
	inorder(root)
	return list
}

// 二叉搜索树的中序遍历序列是有序数组
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	if prev != nil {
		if root.Val == prev.Val {
			count++
		} else {
			count = 1
		}
	} else {
		count = 1
	}
	if count == maxCount {
		list = append(list, root.Val)
	} else if count > maxCount {
		list = append([]int{}, root.Val)
		maxCount = count
	}
	prev = root
	inorder(root.Right)
}
