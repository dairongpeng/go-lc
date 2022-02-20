package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

// GetIntersectNode 给定两个链表的头节点，判断两个链表是否相交，如果两个链表相交，请返回相交的第一个节点，不相交则返回nil
func GetIntersectNode(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	// head1的第一个入环节点
	loop1 := head1.GetLoopNode()
	// head2的第一个入环节点
	loop2 := head2.GetLoopNode()
	// 两个无环链表是否相交的情况
	// 由于每个节点只有一个next指针，则如果两个无环相交则相交之后就只剩下公共部分
	// 方法1把第一条链表放到set中，第二个链表依次查在不在该set中，第一个找到的就是
	// 方法2
	// 把链表1走到尾结点end1，记录长度l1
	// 把链表1走到尾结点end2，记录长度l2
	// 如果end1和end2的内存地址不同一定不相交
	// 如果end1==end2，则(1)长的链表从头结点先走保证和短链表相同长度的位置，再以此往下走，第一次相同节点
	// (2)则依次从尾结点出发，找第一次出现内存地址不相同的那个节点，该节点的next节点就是第一次相交的节点
	if loop1 == nil && loop2 == nil {
		return NoLoop(head1, head2)
	}
	// 一个为空，另外一个不为空不可能相交。两个都不为空的情况下共用一个环
	if loop1 != nil && loop2 != nil {
		return BothLoop(head1, loop1, head2, loop2)
	}
	return nil
}

// GetLoopNode 给定一个链表头节点，找到链表第一个入环节点，如果无环，返回nil
func (head *Node) GetLoopNode() *Node {
	// 少于三个节点，无环
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	idx1 := head.Next      // idx1 为慢指针下标
	idx2 := head.Next.Next // idx2 为快指针下标
	for idx1 != idx2 {
		if idx2.Next == nil || idx2.Next.Next == nil {
			return nil
		}
		idx2 = idx2.Next.Next
		idx1 = idx1.Next
	}
	// idx1 == idx2 相遇。快指针回到开头，满指针停在原地
	idx2 = head
	for idx1 != idx2 {
		// 此时快慢指针每次移动相同步数
		idx1 = idx1.Next
		idx2 = idx2.Next
	}
	// 再次相遇的点，即为成环的点
	return idx1
}

// NoLoop 给定两个无环链表的头节点, 如果相交返回相交节点，如果不想交返回nil
func NoLoop(head1, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1 := head1
	cur2 := head2
	n := 0
	// 遍历链表1，记录节点个数
	for cur1.Next != nil {
		n++
		cur1 = cur1.Next
	}
	// 遍历链表2，记录节点个数
	for cur2.Next != nil {
		n--
		cur2 = cur2.Next
	}
	// 由于是单链表，所以如果两个无环链表相交，必定是从相交的点往后，为公共部分,类似这种形状：
	// ------------------------->
	//       1 2 3 4
	//               5 7 8 9 6 3
	// 2 5 8 2 4 5 7
	if cur1 != cur2 { // 遍历完两个链表，如果有环进入公共部分，所以cur1和cur2必定地址相同
		return nil
	}
	// n绝对值此时为  :  链表1长度减去链表2长度的值
	// 谁长，谁的头变成cur1,谁短，谁的头节点变为cur2
	if n > 0 { // 链表1 长
		cur1 = head1
		cur2 = head2
	} else { // 链表2长或两个链表相等
		cur1 = head2
		cur2 = head1
	}
	// 取n的绝对值
	n = int(math.Abs(float64(n)))
	for n != 0 { // 长链表，先走到差值位置。此时长短链表剩余的长度相等
		n--
		cur1 = cur1.Next
	}

	for cur1 != cur2 { // 两个链表共同向下移动，一旦地址相等，即为第一个相交的节点
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}

// BothLoop 两个有环链表，返回第一个相交节点，如果不想交返回nil
//  head1的入环节点是loop1，head2的入环节点是loop2
func BothLoop(head1, loop1, head2, loop2 *Node) *Node {
	var cur1 *Node = nil
	var cur2 *Node = nil
	// 类似第一种都无环的情况
	// 由于是单链表，那么最终成环的是两个链表的公共部门
	// ---------------------------------------->
	//       1 2 3 4
	//
	//               5 7 8 9 6 3 4 5 7 8
	// 							   3   9
	//								 7
	// 2 5 8 2 4 5 7
	if loop1 == loop2 { // 情况1： 公共部分在成环前
		cur1 = head1
		cur2 = head2
		n := 0
		for cur1 != loop1 {
			n++
			cur1 = cur1.Next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.Next
		}
		// 谁长，谁的头变成cur1,谁短，谁的头节点变为cur2
		if n > 0 { // 链表1 长
			cur1 = head1
			cur2 = head2
		} else { // 链表2长或两个链表相等
			cur1 = head2
			cur2 = head1
		}
		// 取n的绝对值
		n = int(math.Abs(float64(n)))
		for n != 0 { // 长链表，先走到差值位置。此时长短链表剩余的长度相等
			n--
			cur1 = cur1.Next
		}
		for cur1 != cur2 { // 两个链表共同向下移动，一旦地址相等，即为第一个相交的节点
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
		return cur1
	} else { // 情况2： 公共部门在环内。
		// 找第一个成环节点转回自身的过程中遇到loop2, 则相交，否则不相交
		cur1 = loop1.Next
		for cur1 != loop1 { // 链表1沿着成环节点转一圈
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.Next
		}
		return nil
	}
}

func main() {
	// 1->2->3->4->5->6->7->null
	head1 := &Node{Val: 1}
	head1.Next = &Node{Val: 2}
	head1.Next.Next = &Node{Val: 3}
	head1.Next.Next.Next = &Node{Val: 4}
	head1.Next.Next.Next.Next = &Node{Val: 5}
	head1.Next.Next.Next.Next.Next = &Node{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = &Node{Val: 7}

	// 0->9->8->6->7->null
	head2 := &Node{Val: 0}
	head2.Next = &Node{Val: 9}
	head2.Next.Next = &Node{Val: 8}
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next // 8 -> head1(6)
	fmt.Println(GetIntersectNode(head1, head2))

	// 1->2->3->4->5->6->7->4...
	head1 = &Node{Val: 1}
	head1.Next = &Node{Val: 2}
	head1.Next.Next = &Node{Val: 3}
	head1.Next.Next.Next = &Node{Val: 4}
	head1.Next.Next.Next.Next = &Node{Val: 5}
	head1.Next.Next.Next.Next.Next = &Node{Val: 6}
	head1.Next.Next.Next.Next.Next.Next = &Node{Val: 7}
	head1.Next.Next.Next.Next.Next.Next.Next = head1.Next.Next.Next // 7 -> head1(4)

	// 0->9->8->2...
	head2 = &Node{Val: 0}
	head2.Next = &Node{Val: 9}
	head2.Next.Next = &Node{Val: 8}
	head2.Next.Next.Next = head1.Next // 8 -> head1(2)
	fmt.Println(GetIntersectNode(head1, head2))

	// 0->9->8->6->4->5->6..
	head2 = &Node{Val: 0}
	head2.Next = &Node{Val: 9}
	head2.Next.Next = &Node{Val: 8}
	head2.Next.Next.Next = head1.Next.Next.Next.Next.Next // 8 -> head1(6)
	fmt.Println(GetIntersectNode(head1, head2))

}
