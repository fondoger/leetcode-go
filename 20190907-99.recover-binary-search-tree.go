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

**BST性质：BST的中序遍历序列是一个有序数据**

首先将问题简化，如果给的是一个有序数组，然后打乱某两个元素的次序，
如何找到这两个元素？
例如：
1 2 3 5 4 6 7
1 6 3 5 2 6 7
==> 从前往后迭代，第一次发现后者小于前者，就假设后者和前者交换了元素，记录为first和second
==> 如果第二次发现后者小于前者，就把second更新为第二次发现的后者
*/

var pre, first, second *TreeNode
var foundAll bool

func recoverTree(root *TreeNode) {
	pre, first, second, foundAll = nil, nil, nil, false
	search(root)
	first.Val, second.Val = second.Val, first.Val
}

func search(root *TreeNode) {
	if root == nil || foundAll {
		return
	}
	search(root.Left)
	if pre != nil && root.Val < pre.Val {
		if first == nil {
			first = pre
			second = root
		} else {
			second = root
			foundAll = true
		}
	}
	pre = root
	search(root.Right)
}
