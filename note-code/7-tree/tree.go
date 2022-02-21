package main

import "fmt"

//type Node struct {
//	// 二叉树节点上的值
//	Val int
//	// 左孩子
//	Left *Node
//	// 右孩子
//	Right *Node
//}

//// Pre 给定二叉树头节点，先序遍历该二叉树
//func (head *Node) Pre() {
//	if head == nil {
//		return
//	}
//	// 获取头节点，打印该头结点
//	fmt.Println(head.Val)
//	// 递归遍历左子树
//	head.Left.Pre()
//	// 递归遍历右子树
//	head.Right.Pre()
//}
//
//// Mid 给定二叉树头节点，中序遍历该二叉树
//func (head *Node) Mid() {
//	if head == nil {
//		return
//	}
//	// 递归遍历左子树
//	head.Left.Mid()
//	// 获取头节点，打印该头结点
//	fmt.Println(head.Val)
//	// 递归遍历右子树
//	head.Right.Mid()
//}
//
//// Pos 给定二叉树头节点，后序遍历该二叉树
//func (head *Node) Pos() {
//	if head == nil {
//		return
//	}
//	// 递归遍历左子树
//	head.Left.Pos()
//	// 递归遍历右子树
//	head.Right.Pos()
//	// 获取头节点，打印该头结点
//	fmt.Println(head.Val)
//}

// Pre 给定二叉树头节点，非递归先序遍历该二叉树
func (head *Node) Pre() {
	fmt.Println("pre-order: ")
	if head != nil {
		// 简单模拟一个栈
		stack := make([]*Node, 0)
		stack = append(stack, head)
		for len(stack) != 0 {
			// 出栈
			hd := stack[len(stack)-1]
			fmt.Println(hd.Val)
			stack = stack[:len(stack)-1]
			// 右孩子入栈
			if hd.Right != nil {
				stack = append(stack, hd.Right)
			}
			// 左孩子入栈
			if hd.Left != nil {
				stack = append(stack, hd.Left)
			}
		}
	}
	fmt.Println()
}

// Mid 给定二叉树头节点，非递归中序遍历该二叉树
func (head *Node) Mid() {
	fmt.Println("Mid-order:")
	if head != nil {
		hd := head
		// 简单模拟一个栈
		stack := make([]*Node, 0)
		for len(stack) != 0 || hd != nil {
			// 整条左边界依次入栈
			if hd != nil {
				stack = append(stack, hd)
				hd = hd.Left
			} else { // 左边界到头弹出一个打印，来到该节点右节点，再把该节点的左树以此进栈
				hd = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Println(hd.Val)
				hd = hd.Right
			}
		}
	}
	fmt.Println()
}

// Pos 给定二叉树头节点，非递归后序遍历该二叉树
func (head *Node) Pos() {
	fmt.Println("pos-order: ")
	if head != nil {
		hd := head
		// 借助两个辅助栈
		s1 := make([]*Node, 0)
		s2 := make([]*Node, 0)
		s1 = append(s1, hd)
		for len(s1) != 0 {
			// 出栈
			hd = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			s2 = append(s2, hd)
			if hd.Left != nil {
				s1 = append(s1, hd.Left)
			}
			if hd.Right != nil {
				s1 = append(s1, hd.Right)
			}
		}
		for len(s2) != 0 {
			v := s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
			fmt.Println(v.Val)
		}
	}
	fmt.Println()
}

// Level 按层遍历二叉树
func (head *Node) Level() {
	if head == nil {
		return
	}
	hd := head
	// 简单实现一个队列
	queue := make([]*Node, 0)
	// 加入头结点
	queue = append(queue, hd)
	// 队列不为空出队打印，把当前节点的左右孩子加入队列
	for len(queue) != 0 {
		// 弹出队列头部的元素
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

//func main() {
//	head := &Node{Val: 1}
//	head.Left = &Node{Val: 2}
//	head.Right = &Node{Val: 3}
//	head.Left.Left = &Node{Val: 4}
//	head.Left.Right = &Node{Val: 5}
//	head.Right.Left = &Node{Val: 6}
//	head.Right.Right = &Node{Val: 7}
//
//	head.Level()
//}
