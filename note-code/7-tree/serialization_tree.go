package main

import "strconv"

//type Node struct {
//	// 二叉树节点上的值
//	Val int
//	// 左孩子
//	Left *Node
//	// 右孩子
//	Right *Node
//}

// PreSerial 二叉树的先序列化
func (head *Node) PreSerial() []string {
	// 简单实现一个队列
	ans := make([]string, 0)
	// 先序的序列化结果依次放入队列中去
	pres(head, ans)
	return ans
}

func pres(head *Node, ans []string) {
	if head == nil {
		ans = append(ans, "")
	} else {
		ans = append(ans, strconv.Itoa(head.Val))
		pres(head.Left, ans)
		pres(head.Right, ans)
	}
}

// BuildByPreQueue 根据先序序列化的结果，重新构建该二叉树。反序列化
func BuildByPreQueue(preQueue []string) *Node {
	if len(preQueue) == 0 {
		return nil
	}

	return preb(preQueue)
}

func preb(preQueue []string) *Node {
	v := preQueue[0]
	preQueue = preQueue[1:]

	// 如果头节点是空的话，返回空
	if v == "" {
		return nil
	}

	// 否则根据第一个值构建先序的头结点
	if iv, err := strconv.Atoi(v); err == nil {
		head := &Node{Val: iv}
		// 递归建立左树
		head.Left = preb(preQueue)
		// 递归建立右树
		head.Right = preb(preQueue)
		return head
	}
	return nil
}
