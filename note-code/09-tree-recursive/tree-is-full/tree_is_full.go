package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Info 包含高度信息，和节点个数信息
type Info struct {
	Height int
	Nodes  int
}

func (head *Node) IsFull() bool {
	if head == nil {
		return true
	}
	all := process(head)
	// 当前二叉树的高度乘以2 是否等于当前二叉树的节点个数，从而可以判断当前二叉树是否是满二叉树
	return all.Height*2-1 == all.Nodes
}

func process(node *Node) *Info {
	if node == nil {
		return &Info{
			Height: 0,
			Nodes:  0,
		}
	}

	leftInfo := process(node.Left)
	rightInfo := process(node.Right)
	// 当前高度
	height := int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height))) + 1
	// 当前节点为树根的二叉树所有节点数
	nodes := leftInfo.Nodes + rightInfo.Nodes + 1
	return &Info{
		Height: height,
		Nodes:  nodes,
	}
}
