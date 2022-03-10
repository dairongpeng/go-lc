package main

//type Node struct {
//	Val  int
//	Next *Node
//}

// NewLinkedList 初始化一个链表 返回链表的头结点
func NewLinkedList(val int) (head *Node) {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

// MidOrUpMidNode 给定一个链表的头节点
// 1. 奇数长度返回中点, 偶数长度返回上中点
func (head *Node) MidOrUpMidNode() *Node {
	// 该链表没有点，有一个点，有两个点的时候都是返回头结点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	// 链表有3个点或以上
	// 快慢指针，快指针一次走两步，慢指针一次走一步
	// 快指针走完，慢指针在中点位置
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// MidOrDownMidNode 给定一个链表的头节点
// 2、奇数长度返回中点，偶数长度返回中下点
func (head *Node) MidOrDownMidNode() *Node {
	// 该链表没有点，有一个点, 返回头结点
	if head == nil || head.Next == nil {
		return head
	}
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// MidOrUpMidPreNode 给定一个链表的头节点
// 3、奇数长度返回中点前一个，偶数长度返回上中点前一个
func (head *Node) MidOrUpMidPreNode() *Node {
	// 该链表没有点，有一个点, 有两个点， 返回头结点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// MidOrDownMidPreNode 给定一个链表的头节点
// 4、奇数长度返回中点前一个，偶数长度返回下中点前一个
func (head *Node) MidOrDownMidPreNode() *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//func main() {
//	// 0 1
//	// 0 1 2
//	// 0 1 2 3
//	// 0 1 2 3 4
//	// 0 1 2 3 4 5
//	// 0 1 2 3 4 5 6
//	// 0 1 2 3 4 5 6 7
//	// 0 1 2 3 4 5 6 7 8
//	hd := &Node{}
//	hd.Next = &Node{Val: 1}
//	hd.Next.Next = &Node{Val: 2}
//	hd.Next.Next.Next = &Node{Val: 3}
//	hd.Next.Next.Next.Next = &Node{Val: 4}
//	hd.Next.Next.Next.Next.Next = &Node{Val: 5}
//	hd.Next.Next.Next.Next.Next.Next = &Node{Val: 6}
//	hd.Next.Next.Next.Next.Next.Next.Next = &Node{Val: 7}
//	// hd.Next.Next.Next.Next.Next.Next.Next.Next = &Node{Val: 8}
//
//	ans1 := hd.MidOrUpMidNode()
//	fmt.Println(fmt.Sprintf("1.奇数长度返回中点，偶数长度返回上中点: %d", ans1.Val))
//	ans2 := hd.MidOrDownMidNode()
//	fmt.Println(fmt.Sprintf("2.奇数长度返回中点，偶数长度返回中下点: %d", ans2.Val))
//	ans3 := hd.MidOrUpMidPreNode()
//	fmt.Println(fmt.Sprintf("3.奇数长度返回中点前一个，偶数长度返回上中点前一个: %d", ans3.Val))
//	ans4 := hd.MidOrDownMidPreNode()
//	fmt.Println(fmt.Sprintf("4.奇数长度返回中点前一个，偶数长度返回下中点前一个: %d", ans4.Val))
//}
