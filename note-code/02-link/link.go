package main

// Node 单项链表节点结构
type Node struct {
	V    int
	Next *Node
}

// DoubleNode 双向链表节点结构
type DoubleNode struct {
	V    int
	Pre  *DoubleNode
	Next *DoubleNode
}

// ReverseLinkedList 单链表翻转
func ReverseLinkedList(head *Node) *Node {
	var pre *Node
	var next *Node
	for head != nil {
		// 把当前节点的下一个节点保存到next
		next = head.Next
		// 当前节点的指向，改为指向前一个节点
		head.Next = pre
		// pre 节点按原链表方向向下移动
		pre = head
		// head 节点按原链表方向向下移动
		head = next
	}
	// 按照原链表方向移动，当前节点为nil退出循环的时候，那么pre节点就是原链表的最后一个节点，链表被成功翻转。
	// 当前头结点pre返回
	return pre
}

// ReverseDoubleLinkedList 双链表翻转
func ReverseDoubleLinkedList(head *DoubleNode) *DoubleNode {
	var pre *DoubleNode
	var next *DoubleNode
	for head != nil {
		// 保留当前节点的next节点的地址
		next = head.Next
		// 当前节点的下一个节点指pre
		head.Next = pre
		// 当前节点的上一个节点指向原链表当前节点的next节点。
		head.Pre = next
		// pre 节点按原链表方向向下移动
		pre = head
		// head 节点按原链表方向向下移动
		head = next
	}
	return pre
}

// RemoveValue 删除链表中值等于target的节点
func RemoveValue(head *Node, target int) *Node {
	// 处理链表头结点的值即等于target的节点
	for head != nil {
		if head.V != target {
			break
		}
		head = head.Next
	}

	// 1、链表中的节点值全部都等于target
	// 2、原始链表为nil
	if head == nil {
		return head
	}

	// head来到第一个不需要删除的位置
	pre := head
	cur := head
	for cur != nil {
		// 当前节点cur往下，有多少v等于target的节点，就删除多少节点
		if cur.V == target {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		// 当前节点向下滑动
		cur = cur.Next
	}
	return head
}
