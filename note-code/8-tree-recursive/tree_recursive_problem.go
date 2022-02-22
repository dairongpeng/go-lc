package main

import "math"

//type Node struct {
//	Val int
//	Left *Node
//	Right *Node
//}

// BalanceInfo 表示递归过程中需要收集每个节点的信息
type BalanceInfo struct {
	// 当前节点为头的树是不是平衡的
	IsBalanced bool
	// 当前节点为头的树的高度是多少
	Height int
}

// IsBalanced 给定二叉树头节点，判断该二叉树是不是平衡二叉树
func (head *Node) IsBalanced() bool {
	return Process(head).IsBalanced
}

func Process(node *Node) *BalanceInfo {
	if node == nil {
		return &BalanceInfo{
			IsBalanced: true,
			Height:     0,
		}
	}
	// 左子树信息
	leftInfo := Process(node.Left)
	// 右子树信息
	rightInfo := Process(node.Right)
	// 高度等于左右最大高度，加上当前头结点的高度1
	curHeight := int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1
	curIsBalanced := true
	// 左树不平衡或者右树不平衡，或者左右两子树高度差超过1
	// 那么当前节点为头的树，不平衡
	if !leftInfo.IsBalanced || !rightInfo.IsBalanced || math.Abs(float64(leftInfo.Height)-float64(rightInfo.Height)) > 1 {
		curIsBalanced = false
	}
	// 加工出当前节点的信息返回
	return &BalanceInfo{
		IsBalanced: curIsBalanced,
		Height:     curHeight,
	}
}
