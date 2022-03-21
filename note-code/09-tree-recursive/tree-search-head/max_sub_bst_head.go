package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Info 任何子树递归过程中,都返回4个信息
type Info struct {
	// 当前节点为头的二叉树中，最大搜索二叉树的头结点
	MaxSubBSTHead *Node
	// 以当前节点为头节点的树，整体是否是二叉搜索树
	IsAllBST bool
	// 最大的满足二叉搜索树树条件的size
	MaxSubBSTSize int
	// 整棵树的最小值
	Min int
	// 整棵树的最大值
	Max int
}

func (head *Node) MaxSubBSTHead() *Node {
	if head == nil {
		return nil
	}
	return process(head).MaxSubBSTHead
}

func process(node *Node) *Info {
	if node == nil {
		return nil
	}

	// 左子树返回的Info信息
	leftInfo := process(node.Left)
	// 右子树返回的Info信息
	rightInfo := process(node.Right)

	min, max := node.Val, node.Val
	var maxSubBSTHead *Node
	maxSubBSTSize := 0

	// 如果当前节点为头的二叉树不是二叉搜索树，则当前Info信息中isAllBST为false
	isAllBST := false

	// case1：当前节点为头的二叉树中存在的最大搜索二叉树与当前节点node无关的情况
	// 左树不为空，加工min, max, maxSubBSTHead, maxSubBSTSize
	if leftInfo != nil {
		min = int(math.Min(float64(min), float64(leftInfo.Min)))
		max = int(math.Max(float64(max), float64(leftInfo.Max)))
		maxSubBSTHead = leftInfo.MaxSubBSTHead
		maxSubBSTSize = leftInfo.MaxSubBSTSize
	}

	// 右树不为空，加工min, max, maxSubBSTHead, maxSubBSTSize
	if rightInfo != nil {
		min = int(math.Min(float64(min), float64(rightInfo.Min)))
		max = int(math.Max(float64(max), float64(rightInfo.Max)))
		if rightInfo.MaxSubBSTSize > maxSubBSTSize {
			maxSubBSTHead = rightInfo.MaxSubBSTHead
			maxSubBSTSize = rightInfo.MaxSubBSTSize
		}
	}

	// case2: 当前节点为头的二叉树中存在的最大搜索二叉树与当前节点node有关的情况
	// 左树整个是二叉搜索树么
	leftIsAllBST := false
	// 右树整个是二叉搜索树么
	rightIsAllBST := false
	// 左树最大值小于node的值是否
	leftMaxVLessNodeV := false
	// 右树的最小值，大于node的值是否
	rightMinMoreNodeV := false
	if leftInfo == nil {
		leftIsAllBST = true
		leftMaxVLessNodeV = true
	} else {
		leftIsAllBST = leftInfo.IsAllBST
		leftMaxVLessNodeV = leftInfo.Max < node.Val
	}

	if rightInfo == nil {
		rightIsAllBST = true
		rightMinMoreNodeV = true
	} else {
		rightIsAllBST = rightInfo.IsAllBST
		rightMinMoreNodeV = rightInfo.Min > node.Val
	}

	// 如果左树是二叉搜索树，右树也是二叉搜索树，当前节点为树根的左树最大值都比当前值小，当前节点为树根的右树最小值都比当前值大
	// 证明以当前节点node为树根的树，也是一个二叉搜索树。满足case2
	if leftIsAllBST && rightIsAllBST && leftMaxVLessNodeV && rightMinMoreNodeV {
		leftSize := 0
		rightSize := 0
		if leftInfo != nil {
			leftSize = leftInfo.MaxSubBSTSize
		}

		if rightInfo != nil {
			rightSize = rightInfo.MaxSubBSTSize
		}

		maxSubBSTHead = node
		maxSubBSTSize = leftSize + rightSize + 1
		// 当前节点整个是二叉搜索树
		isAllBST = true
	}

	return &Info{
		MaxSubBSTHead: maxSubBSTHead,
		IsAllBST:      isAllBST,
		MaxSubBSTSize: maxSubBSTSize,
		Min:           min,
		Max:           max,
	}

}
