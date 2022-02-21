package main

import "math"

//type Node struct {
//	// 二叉树节点上的值
//	Val int
//	// 左孩子
//	Left *Node
//	// 右孩子
//	Right *Node
//}

// MaxWidthUseMap 给定二叉树头节点，找到该二叉树的最大宽度，借助map结构实现
func (head *Node) MaxWidthUseMap() int {
	if head == nil {
		return 0
	}
	hd := head
	queue := make([]*Node, 0)
	queue = append(queue, hd)

	// map的Key:节点 map的value:节点属于哪一层
	levelMap := make(map[*Node]int, 0)
	// 头节点head属于第一层
	levelMap[hd] = 1
	// 当前正在统计那一层的宽度
	curLevel := 1
	// 当前curLevel层，宽度目前是多少
	curLevelNodes := 0
	// 用来保存所有层的最大宽度的值
	max := 0
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curNodeLevel := levelMap[cur]
		// 当前节点的左孩子不为空，队列加入左孩子，层数在之前层上加1
		if cur.Left != nil {
			levelMap[cur.Left] = curNodeLevel + 1
			queue = append(queue, cur.Left)
		}
		// 当前节点的右孩子不为空，队列加入右孩子，层数也变为当前节点的层数加1
		if cur.Right != nil {
			levelMap[cur.Right] = curNodeLevel + 1
			queue = append(queue, cur.Right)
		}
		// 当前层等于正在统计的层数，不结算
		if curNodeLevel == curLevel {
			curLevelNodes++
		} else {
			// 新的一层，需要结算
			// 得到目前为止的最大宽度
			max = int(math.Max(float64(max), float64(curLevelNodes)))
			curLevel++
			// 结算后，当前层节点数设置为1
			curLevelNodes = 1
		}
	}
	// 由于最后一层，没有新的一层去结算，所以这里单独结算最后一层
	max = int(math.Max(float64(max), float64(curLevelNodes)))
	return max
}

// MaxWidthNoMap 给定二叉树头节点，找到该二叉树的最大宽度，不借助map实现
func (head *Node) MaxWidthNoMap() int {
	if head == nil {
		return 0
	}

	hd := head
	queue := make([]*Node, 0)
	queue = append(queue, hd)

	// 当前层，最右节点是谁，初始head的就是本身
	curEnd := head
	// 如果有下一层，下一层最右节点是谁
	var nextEnd *Node = nil
	// 全局最大宽度
	max := 0
	// 当前层的节点数
	curLevelNodes := 0
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		// 左边不等于空，加入左
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			// 孩子的最右节点暂时为左节点
			nextEnd = cur.Left
		}
		// 右边不等于空，加入右
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			// 如果有右节点，孩子层的最右要更新为右节点
			nextEnd = cur.Right
		}
		// 由于最开始弹出当前节点，那么该层的节点数加一
		curLevelNodes++
		// 当前节点是当前层最右的节点，进行结算
		if cur == curEnd {
			// 当前层的节点和max进行比较，计算当前最大的max
			max = int(math.Max(float64(max), float64(curLevelNodes)))
			// 即将进入下一层，重置下一层节点为0个节点
			curLevelNodes = 0
			// 当前层的最右，直接更新为找出来的下一层最右
			curEnd = nextEnd
		}
	}
	return max
}
