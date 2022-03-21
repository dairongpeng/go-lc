package main

import "fmt"

type Node struct {
	// 二叉树节点上的值
	Val int
	// 左孩子
	Left *Node
	// 右孩子
	Right *Node
	// 指向父亲的指针
	Parent *Node
}

// GetSuccessorNode 给定二叉树的一个节点node，返回该节点的后继节点
func GetSuccessorNode(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.Right != nil {
		return getLeftMost(node.Right)
	} else { // 无右子树
		parent := node.Parent
		// 当前节点是其父亲节点右孩子，继续
		for parent != nil && parent.Right == node {
			node = parent
			parent = node.Parent
		}
		return parent
	}
}

// 找右树上的最左节点
func getLeftMost(node *Node) *Node {
	if node == nil {
		return node
	}
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	head := &Node{Val: 6}
	head.Parent = nil
	head.Left = &Node{Val: 3}
	head.Left.Parent = head
	head.Left.Left = &Node{Val: 1}
	head.Left.Left.Parent = head.Left
	head.Left.Left.Right = &Node{Val: 2}
	head.Left.Left.Right.Parent = head.Left.Left
	head.Left.Right = &Node{Val: 4}
	head.Left.Right.Parent = head.Left
	head.Left.Right.Right = &Node{Val: 5}
	head.Left.Right.Right.Parent = head.Left.Right
	head.Right = &Node{Val: 9}
	head.Right.Parent = head
	head.Right.Left = &Node{Val: 8}
	head.Right.Left.Parent = head.Right
	head.Right.Left.Left = &Node{Val: 7}
	head.Right.Left.Left.Parent = head.Right.Left
	head.Right.Right = &Node{Val: 10}
	head.Right.Right.Parent = head.Right

	test := head.Left.Left
	fmt.Println(fmt.Sprintf("节点：%d的后继节点为%d", test.Val, GetSuccessorNode(test).Val))

	test = head.Left.Left.Right
	fmt.Println(fmt.Sprintf("节点：%d的后继节点为%d", test.Val, GetSuccessorNode(test).Val))

	test = head.Left.Right.Left
	fmt.Println(fmt.Sprintf("节点：%d的后继节点为%d", test.Val, GetSuccessorNode(test).Val))

	test = head
	fmt.Println(fmt.Sprintf("节点：%d的后继节点为%d", test.Val, GetSuccessorNode(test).Val))

	test = head.Right.Left.Left
	fmt.Println(fmt.Sprintf("节点：%d的后继节点为%d", test.Val, GetSuccessorNode(test).Val))
}
