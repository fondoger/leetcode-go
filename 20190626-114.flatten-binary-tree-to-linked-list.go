package main

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 这题本来可以写得非常简单，一个递归即可
// 但是由于golang的测试程序问题，多次调用flatten函数会影响后面的结果
// 因此我利用golang的闭包，把它改成了如下面目全非的样子
/*java版代码
private TreeNode prev = null;

public void flatten(TreeNode root) {
	if (root == null) return;
	flatten(root.Rigth)
	flatten(root.Left)
	root.Right = prev
	root.Left = null
	prev = root
}
*/
// Solution 2: 简单递归
func flatten(root *TreeNode) {
	var prev *TreeNode
	var flat func(*TreeNode)
	flat = func(root *TreeNode) {
		if root == nil {
			return
		}
		flat(root.Right)
		flat(root.Left)
		root.Right = prev
		root.Left = nil
		prev = root
	}
	flat(root)
}

// Solution 1: 我自己想出来的笨办法：递归
func flatten_MY_STUPID_SOLUTION(root *TreeNode) {
	helper114(root)
}

// flattern并返回尾部结点
func helper114(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil {
		return helper114(root.Right)
	} else if root.Right == nil {
		root.Right, root.Left = root.Left, nil
		return helper114(root.Right)
	}
	leftRoot := root.Left
	root.Left = nil
	leftTail := helper114(leftRoot)
	root.Right, leftTail.Right = leftRoot, root.Right
	return helper114(leftTail.Right)
}
