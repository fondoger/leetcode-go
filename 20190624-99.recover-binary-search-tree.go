package main

import "math"

/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* 我喜欢这道题

首先将问题简化，如果给的是一个有序数组，然后打乱某两个元素的次序，
如何找到这两个元素？
例如：
1 2 3 5 4 6 7
==> 从前往后迭代，第一次发现后者小于前者，就假设后者和前者交换了元素，记录为first和second
==> 如果第二次发现后者小于前者，就把second更新为第二次发现的后者
*/

var prev *TreeNode
var first, second *TreeNode
var findAll bool

// 关键：不改变指针，仅交换结点的值
func recoverTree(root *TreeNode) {
	prev = &TreeNode{math.MinInt64, nil, nil}
	first, second = nil, nil
	findAll = false
	helper99(root)
	first.Val, second.Val = second.Val, first.Val
}

func helper99(root *TreeNode) {
	if root == nil || findAll {
		return
	}
	helper99(root.Left)
	if first == nil && root.Val < prev.Val {
		first = prev
		second = root
	} else if first != nil && root.Val < prev.Val {
		second = root
		findAll = true
	}
	prev = root
	helper99(root.Right)
}
