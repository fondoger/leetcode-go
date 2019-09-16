package main

/*
 * @lc app=leetcode id=236 lang=golang
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 *
 */

var foundP, foundQ bool

// 返回p结点或者q结点或者公共结点
// 如果遇到第一个是p结点或者q结点，可以直接返回
// 如果左子树没有任何结点，那么肯定不会在左子树上
// 如果右子树不包含p或者q结点，那么肯定在右子树上
// 如果左右子树分别包含了p或q结点，那么肯定在自身
func lowestCommonAncestor_Solution1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pathP, pathQ []*TreeNode
	pathToNode(&pathP, root, p, 0)
	pathToNode(&pathQ, root, q, 0)
	// 找到各自遍历路径的最后一个公共祖先
	pos := 0
	for pos+1 < len(pathP) && pos+1 < len(pathQ) && pathP[pos+1] == pathQ[pos+1] {
		pos++
	}
	return pathP[pos]
}

func pathToNode(path *[]*TreeNode, root *TreeNode, node *TreeNode, depth int) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root.Val == node.Val ||
		pathToNode(path, root.Left, node, depth+1) ||
		pathToNode(path, root.Right, node, depth+1) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
