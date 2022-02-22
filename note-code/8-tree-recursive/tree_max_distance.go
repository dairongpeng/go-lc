package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type DistanceInfo struct {
	// 当前节点为树根的情况下，该树的最大距离
	MaxDistance int
	// 当前节点为树根的情况下，该树的高度
	Height int
}

// GetMaxDistance 给定二叉树头节点，求该二叉树的最大距离
func (head *Node) GetMaxDistance() int {
	return Process2(head).MaxDistance
}

func Process2(node *Node) *DistanceInfo {
	// base case
	if node == nil {
		return &DistanceInfo{
			MaxDistance: 0,
			Height:      0,
		}
	}

	// 左树信息
	leftInfo := Process2(node.Left)
	// 右树信息
	rightInfo := Process2(node.Right)
	// 用左右树的信息，加工当前节点自身的info
	// 自身的高度是，左右较大的高度加上自身节点高度1
	curHeight := int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height)))
	// 自身最大距离，(左右树最大距离)和(左右树高度相加再加1)，求最大值
	curMaxDistance := int(math.Max(
		math.Max(float64(leftInfo.MaxDistance), float64(rightInfo.MaxDistance)),
		float64(leftInfo.Height+rightInfo.Height+1)))
	// 自身的info返回
	return &DistanceInfo{
		MaxDistance: curMaxDistance,
		Height:      curHeight,
	}
}
