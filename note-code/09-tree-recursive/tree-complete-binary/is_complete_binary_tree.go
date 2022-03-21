package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsCBTWithProcess 递归二叉树判断一颗二叉树是否是完全二叉树
func (head *Node) IsCBTWithProcess() bool {
	if head == nil {
		return true
	}
	return process(head).IsCBT
}

type Info struct {
	IsFull bool
	IsCBT  bool
	Height int
}

func process(node *Node) *Info {
	// 如果是空树，我们封装Info而不是返回为空
	// 方便下文不需要额外增加判空处理
	if node == nil {
		return &Info{
			IsFull: true,
			IsCBT:  true,
			Height: 0,
		}
	}

	leftInfo := process(node.Left)
	rightInfo := process(node.Right)

	// 整合当前节点的Info
	// 高度信息=左右树最大高度值+1
	height := int(math.Max(float64(leftInfo.Height), float64(rightInfo.Height)))
	// node是否是满二叉树信息=左右都是满且左右高度一样
	isFull := leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height
	isBST := false
	if isFull { // 满二叉树是完全二叉树
		isBST = true
	} else { // 以node为头整棵树，不满
		// 左右都是完全二叉树才有讨论的必要
		if leftInfo.IsCBT && rightInfo.IsCBT {
			// 第二种情况。左树是完全二叉树，右树是满二叉树，左树高度比右树高度大1
			if leftInfo.IsCBT && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 {
				isBST = true
			}
			// 第三种情况。左树满，右树满，且左树高度比右树高度大1
			if leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 {
				isBST = true
			}
			// 第四种情况。左树满，右树是完全二叉树，且左右树高度相同
			if leftInfo.IsFull && rightInfo.IsCBT && leftInfo.Height == rightInfo.Height {
				isBST = true
			}
		}
	}
	return &Info{
		IsFull: isFull,
		IsCBT:  isBST,
		Height: height,
	}
}

// IsCBTWithWidth 宽度优先遍历判断一颗二叉树是否是完全二叉树
func (head *Node) IsCBTWithWidth() bool {
	if head == nil {
		return true
	}

	hd := head

	var queue = make([]*Node, 0)
	// 是否遇到过左右两个孩子不双全的节点
	leaf := false
	var l *Node = nil
	var r *Node = nil
	queue = append(queue, hd)
	for len(queue) != 0 {
		hd = queue[0]
		queue = queue[1:]
		l = hd.Left
		r = hd.Right
		// 如果遇到了不双全的节点之后，又发现当前节点不是叶节点
		if leaf && (l != nil || r != nil) || (l == nil && r == nil) {
			return false
		}
		if l != nil {
			queue = append(queue, l)
		}
		if r != nil {
			queue = append(queue, r)
		}
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}
