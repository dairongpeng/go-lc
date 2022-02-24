package main

import "math"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Info 任何子树递归过程中,都返回4个信息
type Info struct {
	// 以当前节点为头节点的树，整体是否是二叉搜索树
	IsAllBST bool
	// 最大的满足二叉搜索树树条件的size
	MaxSubBSTSize int
	// 整棵树的最小值
	Min int
	// 整棵树的最大值
	Max int
}

// MaxSubBSTSize 给定二叉树头节点，返回该二叉树的最大二叉搜索子树的大小
func (head *Node) MaxSubBSTSize() int {
	if head == nil {
		return 0
	}
	return process(head).MaxSubBSTSize
}

func process(node *Node) *Info {
	if node == nil {
		return nil
	}

	// 左子树返回的Info信息
	leftInfo := process(node.Left)
	// 右子树返回的Info信息
	rightInfo := process(node.Right)

	// 加工我自身的info
	min, max := node.Val, node.Val
	// 左树不为空，加工min和max
	if leftInfo != nil {
		min = int(math.Min(float64(min), float64(leftInfo.Min)))
		max = int(math.Max(float64(max), float64(leftInfo.Max)))
	}

	// 右树不为空，加工min和max
	if rightInfo != nil {
		min = int(math.Min(float64(min), float64(rightInfo.Min)))
		max = int(math.Max(float64(max), float64(rightInfo.Max)))
	}

	// case1: 与node无关的情况。当前二叉树存在的最大搜索二叉树的最大大小，是左右树存在的最大二叉搜索树的较大的
	maxSubBSTSize := 0
	if leftInfo != nil {
		maxSubBSTSize = leftInfo.MaxSubBSTSize
	}
	if rightInfo != nil {
		maxSubBSTSize = int(math.Max(float64(maxSubBSTSize), float64(rightInfo.MaxSubBSTSize)))
	}
	// 如果当前节点为头的二叉树不是二叉搜索树，则当前Info信息中isAllBST为false
	isAllBST := false

	// case2：与node有关的情况
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

		// 当前节点为树根的二叉搜索树的节点大小是左树存在的最大二叉搜索树的大小，加上右树存在的最大的二叉搜索树的大小，加上当前node节点1
		maxSubBSTSize = leftSize + rightSize + 1
		// 当前节点整个是二叉搜索树
		isAllBST = true
	}

	return &Info{
		IsAllBST:      isAllBST,
		MaxSubBSTSize: maxSubBSTSize,
		Min:           min,
		Max:           max,
	}

}
