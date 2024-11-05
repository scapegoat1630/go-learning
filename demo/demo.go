package main

import "fmt"

// 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
// 3
// 9 | 15,20,7

// 20,15,7  15,20,7

// 20
// 15 | 7

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	idx := 0
	for k, v := range inorder {
		if v == root.Val {
			idx = k
			break
		}
	}
	if idx == 0 {
		root.Right = buildTree(preorder[1:], inorder[1:])
		return root
	}
	if idx == len(inorder)-1 {
		root.Left = buildTree(preorder[1:], inorder[0:len(inorder)-1])
		return root
	}
	inorderLeft := inorder[0:idx]
	inorderRight := inorder[idx+1:]
	preOrderLeft := preorder[1 : idx+1]
	preOrderRight := preorder[idx+1:]
	root.Left = buildTree(preOrderLeft, inorderLeft)
	root.Right = buildTree(preOrderRight, inorderRight)
	return root
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	PrintTree(root)
}
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Printf("null ")
		return
	}
	PrintTree(root.Left)
	fmt.Printf("%d ", root.Val)
	PrintTree(root.Right)
}
