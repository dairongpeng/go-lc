package main

// 判断一个链表是否是回文链表
// 解法1：遍历链表，把链表放入一个数组中，倒序遍历数组切片，和正序遍历数组切片依次对比，都相等即为回文，实现略
// 解法2：遍历链表，按照快慢指针使慢指针定位到链表的中点位置，依次把慢指针指向的值写入数组切片中，此时数组中保存着链表前一半的元素。
//			且slow的位置是链表的中间位置。按倒序遍历数组切片，与按slow位置顺序遍历链表的元素依次对比，如果都相等，则为回文。实现略
// 解法2比解法1省了一半的空间，解法一空间复杂度O(n)，解法2空间复杂度为O(n/2)
// 解法3：不使用额外空间，空间复杂度为O(1)

//type Node struct {
//	Val  int
//	Next *Node
//}

// IsPalindrome 给定链表头节点，判断该链表是不是回文链表。空间复杂度为O(1)
func (head *Node) IsPalindrome() bool {
	// 链表为空，或者链表只有一个节点，是回文结构
	if head == nil || head.Next == nil {
		return true
	}
	// 慢指针
	slow := head
	// 快指针
	fast := head
	for fast.Next != nil && fast.Next.Next != nil { // 循环结束，slow停在链表中点位置
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next // 快指针回到中点的下一个节点，之后快指针将从每次走两步，变为每次走一步
	slow.Next = nil  // 从中点截断链表，mid.next -> nil
	var tmp *Node
	// 对原链表右半部分，进行逆序，逆序后，从原尾节点指向中点
	for fast != nil {
		tmp = fast.Next  // tmp暂时保存fast的下一个节点
		fast.Next = slow // 翻转链表指向
		slow = fast      // slow 移动
		fast = tmp       // fast 移动
	}

	// tmp指针记录最后的位置，之后把右半部再逆序回原来的次序
	tmp = slow
	fast = head
	var res = true
	for slow != nil && fast != nil { // 原链表的左右部门进行回文对比
		if slow.Val != fast.Val {
			res = false
			break
		}
		slow = slow.Next // 从原链表头节点，往原链表中间节点移动
		fast = fast.Next // 从原链表尾节点，往原链表中间节点移动
	}
	slow = tmp.Next
	tmp.Next = nil
	// 把原链表右半部分再逆序回来
	for slow != nil {
		fast = slow.Next
		slow.Next = tmp
		tmp = slow
		slow = fast
	}
	// 返回回文的判断结果 true
	return res
}
