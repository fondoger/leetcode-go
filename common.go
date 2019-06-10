package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入："[3,2,0,-4]"
// 输出：ListNode单向链表
func buildLinkedList(str string) *ListNode {
	str = strings.Replace(str, " ", "", -1)
	str = str[1 : len(str)-1] // 去除[]中括号
	if str == "" {
		return nil
	}
	items := strings.Split(str, ",")
	dummy := new(ListNode)
	cur := dummy
	for _, item := range items {
		value, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		cur.Next = &ListNode{value, nil}
		cur = cur.Next
	}
	return dummy.Next
}

// 打印 ListNode 单向链表
func printLinkedList(head *ListNode) {
	print("List: [")
	cur := head
	count := 0
	for cur != nil {
		if cur != head {
			print(",")
		}
		print(cur.Val)
		cur = cur.Next
		count++
	}
	print("]", ", size=", count)
	println()
}

// 调用元素默认的toString()方法
func printIntArray(arr []int) {
	print("Array: [")
	for _, item := range arr {
		print(item, " ")
	}
	print("]", ", length=", len(arr))
	println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入二叉树：
//			0
//		  /	  \
//		 -3	   9
//		/	  /
//	  -10    5
// 输出层序遍历结果:[0,-3,9,-10,null,5]
// 假设树是一个满二叉树，层序遍历
func serializeBinaryTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	res := "[" + strconv.Itoa(root.Val)
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root.Left, root.Right)
	nilCount := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			nilCount++
		} else {
			for nilCount > 0 {
				res += ",null"
				nilCount--
			}
			res += "," + strconv.Itoa(node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	res += "]"
	return res
}

func strToTreeNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return &TreeNode{value, nil, nil}
}

func deserializeBinaryTree(str string) *TreeNode {
	str = str[1 : len(str)-1] // 去除[]中括号
	items := strings.Split(str, ",")
	root := strToTreeNode(items[0])
	if root == nil {
		return nil
	}
	items = items[1:]
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)
	for len(queue) > 0 && len(items) > 0 {
		node := queue[0]
		queue = queue[1:]
		if len(items) > 0 {
			node.Left = strToTreeNode(items[0])
			items = items[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		if len(items) > 0 {
			node.Right = strToTreeNode(items[0])
			items = items[1:]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 从LeetCode形式的二叉树结点数组构建二叉树
func buildBinaryTree(str string) *TreeNode {
	return deserializeBinaryTree(str)
}

// 按LeetCode方式层序打印二叉树
func printBinaryTree(root *TreeNode) {
	println(serializeBinaryTree(root))
}
