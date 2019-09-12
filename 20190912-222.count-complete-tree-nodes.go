package main

/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Solution 2: 如果遇到完美二叉树就立即返回，否则递归查找
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLeft(root)
	right := countRight(root)
	if left == right {
		return left - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func countLeft(root *TreeNode) int {
	res := 1
	for root != nil {
		root = root.Left
		res = res << 1
	}
	return res
}

func countRight(root *TreeNode) int {
	res := 1
	for root != nil {
		root = root.Right
		res = res << 1
	}
	return res
}

// Solution 1: 完全二叉树(complete tree)的左子树或右子树必有一个是满二叉树(full tree)
func countNodes_Solution2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	// 检查子树是否为满二叉树，如果是的话，只需要计算右子树
	// 如果不是的话，只需要计算左子树
	left, right := root.Left, root.Left
	for right != nil {
		left = left.Left
		right = right.Right
		count = count << 1
	}
	// right == nil
	if left == nil { // 左子树是满二叉树
		return count + countNodes(root.Right)
	} else { //左子树不是满二叉树，右子树必为满二叉树
		return count + countNodes(root.Left)
	}
}
