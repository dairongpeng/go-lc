package main

type CopyLinkedListNode struct {
	Val  int
	Next *CopyLinkedListNode
	Rand *CopyLinkedListNode
}

// 方法1
func (head *CopyLinkedListNode) copyListWithRand1() *CopyLinkedListNode {
	m := make(map[*CopyLinkedListNode]*CopyLinkedListNode)
	cur := head
	for cur != nil {
		// 当前节点，克隆出来一个相同值的新节点加入map中
		m[cur] = &CopyLinkedListNode{Val: cur.Val}
		cur = cur.Next
	}
	// 当前节点从头开始
	cur = head
	for cur != nil {
		// cur原节点，m[cur]是新节点
		m[cur].Next = m[cur.Next]
		m[cur].Rand = m[cur.Rand]
		cur = cur.Next
	}
	// 返回原头结点的映射，也就是克隆链表的头结点
	return m[head]
}

// 方法2
func (head *CopyLinkedListNode) copyListWithRand2() *CopyLinkedListNode {
	if head == nil {
		return head
	}
	cur := head
	var next *CopyLinkedListNode = nil
	// 克隆出来的node放在原本node和next指向的node中间
	// 1 -> 2
	// 1 -> 1' -> 2
	for cur != nil {
		// cur 老节点   next 老的下一个节点
		next = cur.Next
		cur.Next = &CopyLinkedListNode{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}
	cur = head
	var curCopy *CopyLinkedListNode = nil
	// set copy node rand
	// 1 -> 1' -> 2 -> 2'
	// 设置新的克隆节点间的rand节点
	for cur != nil {
		// cur 老节点
		// cur.next => 新的 copy出来的节点
		next = cur.Next.Next
		curCopy = cur.Next
		if cur.Rand != nil {
			curCopy.Rand = cur.Rand.Next
		} else {
			curCopy.Rand = nil
		}
		cur = next
	}
	// 老的头结点：head 新克隆出来的头结点： head.next
	res := head.Next
	cur = head
	// split，分离原本节点组成的链表和克隆节点组成的链表
	for cur != nil {
		next = cur.Next.Next
		curCopy = cur.Next
		cur.Next = next
		if next != nil {
			curCopy.Next = next.Next
		} else {
			curCopy.Next = nil
		}
		cur = next
	}
	return res
}
